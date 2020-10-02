package api

import (
	"orderContext/application/query"
	"orderContext/infrastructure"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
)

const orderBaseUrl string = "/orders"
const version string = "v1"

func RegisterHandlers(e *echo.Echo) {

	v1 := e.Group("/api/" + version)
	{
		repository := infrastructure.NewOrderRepository()
		service := query.NewOrderQueryService(repository)
		eventBus := infrastructure.NewNoBus()
		timeout := viper.GetInt("context.timeout")

		commandController := newOrderCommandController(repository, eventBus, timeout)
		queryController := newOrderQueryController(service)

		v1.GET(orderBaseUrl, queryController.getOrders)
		v1.GET(orderBaseUrl+"/:id", queryController.getOrder)

		v1.POST(orderBaseUrl, commandController.create)

		v1.PUT(orderBaseUrl+"/pay"+"/:id", commandController.pay)
		v1.PUT(orderBaseUrl+"/ship"+"/:id", commandController.ship)

	}

}
