package main

import (
	_ "ordercontext/docs"

	"ordercontext/internal/api"

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

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api.RegisterHandlers(e, cfg)

	e.Logger.Fatal(e.Start(cfg.Server.Address))
}

func must(fn func() error) {
	err := fn()
	if err != nil {
		panic(err)
	}
}
