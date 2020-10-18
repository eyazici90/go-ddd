package api

import (
	"time"

	"github.com/labstack/echo/v4"

	"orderContext/application/query"
	"orderContext/infrastructure"
	"orderContext/infrastructure/store"
)

const orderBaseUrl string = "/orders"
const version string = "v1"

func RegisterHandlers(e *echo.Echo, cfg Config) {

	v1 := e.Group("/api/" + version)
	{
		mStore := store.NewMongoStore(cfg.MongoDb.Url, cfg.MongoDb.Database, time.Duration(cfg.Context.Timeout)*time.Second)

		repository := infrastructure.NewOrderMongoRepository(mStore)

		service := query.NewOrderQueryService(repository)
		eventBus := infrastructure.NewNoBus()

		commandController := newOrderCommandController(repository, eventBus, cfg.Context.Timeout)
		queryController := newOrderQueryController(service)

		v1.GET(orderBaseUrl, queryController.getOrders)
		v1.GET(orderBaseUrl+"/:id", queryController.getOrder)

		v1.POST(orderBaseUrl, commandController.create)

		v1.PUT(orderBaseUrl+"/pay"+"/:id", commandController.pay)
		v1.PUT(orderBaseUrl+"/ship"+"/:id", commandController.ship)

	}

}