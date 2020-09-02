package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "order context api  is running!")
	})

	controller := NewOrderController()

	e.GET("/order", controller.GetOrders)

	e.GET("/order"+"/:id", controller.GetOrder)

	e.PUT("/order"+"/pay"+"/:id", controller.Pay)

	e.POST("/order", controller.Create)

}
