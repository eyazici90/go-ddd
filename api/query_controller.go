package api

import (
	"context"
	"orderContext/application/query"

	"github.com/labstack/echo/v4"
)

type orderQueryController struct {
	orderservice query.OrderQueryService
}

func newOrderQueryController(s query.OrderQueryService) orderQueryController {
	return orderQueryController{
		orderservice: s,
	}
}

// GetOrder godoc
// @Summary Get orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} order.Order
// @Router /order [get]
func (o *orderQueryController) getOrders(c echo.Context) error {
	return get(c, o.orderservice.GetOrders(c.Request().Context()))
}

// GetOrder godoc
// @Summary Get order
// @Description Get order
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} order.Order
// @Param id path string true "id"
// @Router /order/{id} [get]
func (o *orderQueryController) getOrder(c echo.Context) error {
	return getById(c, func(ctx context.Context, id string) interface{} {
		return o.orderservice.GetOrder(ctx, id)
	})
}
