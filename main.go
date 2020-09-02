package main

import (
	"net/http"
	"orderContext/api"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "order context api  is running!")
	})

	controller := api.NewOrderController()

	e.GET("/order", controller.GetOrders)

	e.PUT("/order"+"/pay"+"/:id", controller.Pay)

	e.POST("/order", controller.Create)

	e.Logger.Fatal(e.Start(":8080"))

}
