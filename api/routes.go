package api

import (
	"time"

	"github.com/labstack/echo/v4"

	"ordercontext/application/query"
	"ordercontext/infrastructure"
	"ordercontext/infrastructure/store/order"
)

const orderBaseURL string = "/orders"
const version string = "v1"

func RegisterHandlers(e *echo.Echo, cfg Config) {

	v1 := e.Group("/api/" + version)
	{
		mStore := infrastructure.NewMongoStore(cfg.MongoDb.URL, cfg.MongoDb.Database, time.Duration(cfg.Context.Timeout)*time.Second)

		repository := order.NewMongoRepository(mStore)

		service := query.NewOrderQueryService(repository)
		eventBus := infrastructure.NewNoBus()

		commandController := newOrderCommandController(repository, eventBus, cfg.Context.Timeout)
		queryController := newOrderQueryController(service)

		v1.GET(orderBaseURL, queryController.getOrders)
		v1.GET(orderBaseURL+"/:id", queryController.getOrder)

		v1.POST(orderBaseURL, commandController.create)

		v1.PUT(orderBaseURL+"/pay"+"/:id", commandController.pay)
		v1.PUT(orderBaseURL+"/ship"+"/:id", commandController.ship)

	}

}
