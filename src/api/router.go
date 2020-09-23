package api

import (
	"orderContext/application/query"
	"orderContext/infrastructure"

	"github.com/labstack/echo/v4"
)

const orderBaseUrl string = "/order"
const version string = "v1"

func RegisterHandlers(e *echo.Echo) {

	v1 := e.Group("/api/" + version)
	{
		r := infrastructure.NewOrderRepository()
		s := query.NewOrderQueryService(r)
		e := infrastructure.NewNoBus()

		commandController := newOrderCommandController(r, e)
		queryController := newOrderQueryController(s)

		v1.GET(orderBaseUrl, queryController.getOrders)
		v1.GET(orderBaseUrl+"/:id", queryController.getOrder)

		v1.POST(orderBaseUrl, commandController.create)

		v1.PUT(orderBaseUrl+"/pay"+"/:id", commandController.pay)
		v1.PUT(orderBaseUrl+"/ship"+"/:id", commandController.ship)

	}

}
