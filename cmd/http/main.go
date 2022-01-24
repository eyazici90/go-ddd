package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	nethttp "net/http"
	"os"
	"time"

	http2 "github.com/eyazici90/go-ddd/internal/http"

	"github.com/eyazici90/go-ddd/internal/app/query"
	"github.com/eyazici90/go-ddd/internal/infra"
	"github.com/eyazici90/go-ddd/internal/infra/inmem"
	"github.com/eyazici90/go-ddd/pkg/must"
	"github.com/eyazici90/go-ddd/pkg/shutdown"

	_ "github.com/eyazici90/go-ddd/docs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Order Application
// @description order context
// @version 1.0
// @host localhost:8080
// @BasePath /http/v1
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
	server := buildServer(w)

	go func() {
		if err := server.Start(); err != nil && err != nethttp.ErrServerClosed {
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

func buildServer(w io.Writer) *http2.Server {
	var cfg http2.Config
	readConfig(&cfg)

	repository := inmem.NewOrderRepository()
	service := query.NewOrderQueryService(repository)
	eventBus := infra.NewNoBus()

	commandController, err := http2.NewCommandController(repository, eventBus, time.Second*time.Duration(cfg.Context.Timeout))
	must.NotFail(err)

	queryController := http2.NewQueryController(service)

	e := echo.New()
	e.Logger.SetOutput(w)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return http2.NewServer(cfg, e, commandController, queryController)
}

func readConfig(cfg *http2.Config) {
	viper.SetConfigFile(`./config.json`)

	must.NotFailF(viper.ReadInConfig)
	must.NotFail(viper.Unmarshal(cfg))
}
