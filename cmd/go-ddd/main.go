package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
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

func run(wr io.Writer) (cleanup func(), err error) {
	defer func() {
		if rvr := recover(); rvr != nil {
			switch r := rvr.(type) {
			case error:
				err = fmt.Errorf("recovery: %w", r)
			default:
				err = fmt.Errorf("recovery: %v", r)
			}
		}
	}()
	setUpSlog(wr)
	server, err := buildServer(wr)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, gohttp.ErrServerClosed) {
			server.Fatal(errors.New("server could not be started"))
		}
	}()

	cleanup = func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(server.Config().Context.Timeout)*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Fatal(err)
		}
	}
	return cleanup, nil
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
	eventBus := infra.NewNoBus()
	mem := inmem.NewOrderRepository()
	qsvc := query.NewService(mem)
	queryCtrl := http.NewQueryController(qsvc)
	cmdCtrl, err := http.NewCommandController(mem, eventBus, time.Second*time.Duration(cfg.Context.Timeout))
	if err != nil {
		return nil, err
	}

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

func setUpSlog(wr io.Writer) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	h := slog.NewTextHandler(wr, opts)
	sl := slog.New(h)
	slog.SetDefault(sl)
}
