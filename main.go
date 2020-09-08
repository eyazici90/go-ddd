package main

import (
	"net/http"
	"orderContext/api"

	"github.com/labstack/echo/v4"

	_ "orderContext/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

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

	api.RegisterHandlers(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8080"))

}
