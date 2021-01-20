package main

import (
	"fmt"
	_ "ordercontext/docs"
	"os"

	"ordercontext/internal/api"
	"ordercontext/internal/application/query"
	"ordercontext/internal/infrastructure"
	"ordercontext/internal/infrastructure/store/order"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var cfg api.Config

func init() {
	viper.SetConfigFile(`config.json`)

	must(viper.ReadInConfig)

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
}

// @title Order Application
// @description order context
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	err, cleanup := run()
	defer cleanup()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func run() (error, func()) {
	repository := order.InMemoryRepository

	service := query.NewOrderQueryService(repository)
	eventBus := infrastructure.NewNoBus()

	commandController := api.NewOrderCommandController(repository, eventBus, cfg.Context.Timeout)
	queryController := api.NewOrderQueryController(service)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	app := api.NewApp(cfg, e, commandController, queryController)

	err := app.Start()

	return err, func() {
		e.Close()
	}
}

func must(fn func() error) {
	err := fn()
	if err != nil {
		panic(err)
	}
}
