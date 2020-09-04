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
		handler := NewOrderHandler()

		v1.GET("/order", handler.GetOrders)

		v1.GET("/order"+"/:id", handler.GetOrder)

		v1.PUT("/order"+"/pay"+"/:id", handler.Pay)

		v1.POST("/order", handler.Create)
	}

}
