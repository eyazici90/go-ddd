package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"ordercontext/internal/api"
	"ordercontext/internal/application/query"
	"ordercontext/internal/infrastructure"
	"ordercontext/internal/infrastructure/store/order"
	"ordercontext/pkg/must"
	"ordercontext/pkg/shutdown"

	_ "ordercontext/docs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var cfg api.Config

func init() {
	viper.SetConfigFile(`./config.json`)

	must.NotFailF(viper.ReadInConfig)
	must.NotFail(viper.Unmarshal(&cfg))
}

// @title Order Application
// @description order context
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cleanup, err := run()
	defer cleanup()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	shutdown.Gracefully()
}

func run() (func(), error) {
	server := buildServer()

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			server.Fatal(errors.New("server could not be started"))
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Context.Timeout)*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Fatal(err)
		}
	}, nil
}

func buildServer() *api.Server {
	repository := order.NewInMemoryRepository()

	service := query.NewOrderQueryService(repository)
	eventBus := infrastructure.NewNoBus()

	commandController := api.NewOrderCommandController(repository, eventBus, time.Second*time.Duration(cfg.Context.Timeout))
	queryController := api.NewOrderQueryController(service)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return api.NewServer(cfg, e, commandController, queryController)
}
