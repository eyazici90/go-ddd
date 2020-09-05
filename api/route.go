package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const orderBaseUrl string = "/order"

func RegisterHandlers(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})

	v1 := e.Group("/api/v1")
	{
		handler := newOrderHandler()

		v1.GET(orderBaseUrl, handler.getOrders)

		v1.GET(orderBaseUrl+"/:id", handler.getOrder)

		v1.PUT(orderBaseUrl+"/pay"+"/:id", handler.pay)

		v1.POST(orderBaseUrl, handler.create)
	}

}
