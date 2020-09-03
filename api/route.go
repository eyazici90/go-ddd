package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "order context api  is running!")
	})

	v1 := e.Group("/api/v1")
	{
		controller := newOrderController()

		v1.GET("/order", controller.getOrders)

		v1.GET("/order"+"/:id", controller.getOrder)

		v1.PUT("/order"+"/pay"+"/:id", controller.pay)

		v1.POST("/order", controller.create)
	}

}
