package main

import (
	"net/http"
	"orderContext/api"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_ "orderContext/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
