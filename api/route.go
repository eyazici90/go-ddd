package api

import (
	"github.com/labstack/echo/v4"
)

const orderBaseUrl string = "/order"
const version string = "v1"

func RegisterHandlers(e *echo.Echo) {

	v1 := e.Group("/api/" + version)
	{
		handler := newOrderHandler()

		v1.GET(orderBaseUrl, handler.getOrders)

		v1.GET(orderBaseUrl+"/:id", handler.getOrder)

		v1.PUT(orderBaseUrl+"/pay"+"/:id", handler.pay)

		v1.POST(orderBaseUrl, handler.create)
	}

}
