package main

import (
	"orderContext/api"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_ "orderContext/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)

	must(viper.ReadInConfig)
}

// @title Order Application
// @description order context
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {

	e := echo.New()

	e.GET("/", api.Health())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}

func must(fn func() error) {
	err := fn()
	if err != nil {
		panic(err)
	}
}
