package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	gohttp "net/http"
	"os"
	"time"

	_ "github.com/eyazici90/go-ddd/docs"
	"github.com/eyazici90/go-ddd/internal/app/query"
	"github.com/eyazici90/go-ddd/internal/http"
	"github.com/eyazici90/go-ddd/internal/infra"
	"github.com/eyazici90/go-ddd/internal/infra/inmem"
	"github.com/eyazici90/go-ddd/pkg/otel"
	"github.com/eyazici90/go-ddd/pkg/shutdown"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Order Application
// @description order context
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	cleanup, err := run(os.Stdout)
	defer cleanup()

	if err != nil {
		fmt.Printf("%v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()
}

func run(w io.Writer) (func(), error) {
	server, err := buildServer(w)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, gohttp.ErrServerClosed) {
			server.Fatal(errors.New("server could not be started"))
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(server.Config().Context.Timeout)*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Fatal(err)
		}
	}, nil
}

func buildServer(wr io.Writer) (*http.Server, error) {
	var cfg http.Config
	if err := readConfig(&cfg); err != nil {
		return nil, err
	}

	cleanup, err := otel.New(context.Background(), &otel.Config{
		Name:    "github.com/eyazici90/go-ddd",
		Version: "1.0.0",
	})
	if err != nil {
		return nil, errors.Join(err, cleanup(context.Background()))
	}
	repo := inmem.NewOrderRepository()
	svc := query.NewOrderQueryService(repo)
	eventBus := infra.NewNoBus()

	cmdCtrl, err := http.NewCommandController(repo, eventBus, time.Second*time.Duration(cfg.Context.Timeout))
	if err != nil {
		return nil, err
	}

	queryCtrl := http.NewQueryController(svc)

	e := echo.New()
	e.Logger.SetOutput(wr)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return http.NewServer(cfg, e, cmdCtrl, queryCtrl), nil
}

func readConfig(cfg *http.Config) error {
	viper.SetConfigFile(`./config.json`)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read cfg: %w", err)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unmarshal cfg: %w", err)
	}
	return nil
}
