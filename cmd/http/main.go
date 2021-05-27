package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ordercontext/internal/api"
	"ordercontext/internal/application/query"
	"ordercontext/internal/infrastructure"
	"ordercontext/internal/infrastructure/store/order"
	"ordercontext/pkg/must"

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
	shutdown, err := run()
	defer shutdown()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func run() (func(), error) {
	server := buildServer()

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			server.Echo().Logger.Fatal("shutting down the server")
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Echo().Logger.Fatal(err)
		}

	}, nil
}

func buildServer() *api.Server {
	repository := order.NewInMemoryRepository()

	service := query.NewOrderQueryService(repository)
	eventBus := infrastructure.NewNoBus()

	commandController := api.NewOrderCommandController(repository, eventBus, cfg.Context.Timeout)
	queryController := api.NewOrderQueryController(service)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return api.NewServer(cfg, e, commandController, queryController)
}
