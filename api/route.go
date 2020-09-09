package api

import (
	"github.com/labstack/echo/v4"
)

const orderBaseUrl string = "/order"
const version string = "v1"

func RegisterHandlers(e *echo.Echo) {

	v1 := e.Group("/api/" + version)
	{
		commandHandler := newOrderCommandHandler()
		queryHandler := newOrderQueryHandler()

		v1.GET(orderBaseUrl, queryHandler.getOrders)

		v1.GET(orderBaseUrl+"/:id", queryHandler.getOrder)

		v1.PUT(orderBaseUrl+"/pay"+"/:id", commandHandler.pay)

		v1.POST(orderBaseUrl, commandHandler.create)
	}

}
