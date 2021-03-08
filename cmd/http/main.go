package main

import (
	"fmt"
	_ "ordercontext/docs"
	"ordercontext/pkg/must"
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
	viper.SetConfigFile(`./config.json`)

	must.NotFailF(viper.ReadInConfig)

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
	cleanup, err := run()
	defer cleanup()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func run() (func(), error) {
	repository := order.NewInMemoryRepository()

	service := query.NewOrderQueryService(repository)
	eventBus := infrastructure.NewNoBus()

	commandController := api.NewOrderCommandController(repository, eventBus, cfg.Context.Timeout)
	queryController := api.NewOrderQueryController(service)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	server := api.NewServer(cfg, e, commandController, queryController)

	err := server.Start()

	return func() {
		e.Close()
	}, err
}
