package api

import (
	"context"

	"ordercontext/internal/application/query"

	"github.com/labstack/echo/v4"
)

type OrderQueryController struct {
	orderservice query.OrderQueryService
}

func NewOrderQueryController(s query.OrderQueryService) *OrderQueryController {
	return &OrderQueryController{
		orderservice: s,
	}
}

// GetOrder godoc
// @Summary Get orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} query.GetOrdersDto
// @Router /orders [get]
func (o *OrderQueryController) getOrders(c echo.Context) error {
	return get(c, o.orderservice.GetOrders(c.Request().Context()))
}

// GetOrder godoc
// @Summary Get order
// @Description Get order
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} query.GetOrderDto
// @Param id path string true "id"
// @Router /orders/{id} [get]
func (o *OrderQueryController) getOrder(c echo.Context) error {
	return getByID(c, func(ctx context.Context, id string) interface{} {
		return o.orderservice.GetOrder(ctx, id)
	})
}
