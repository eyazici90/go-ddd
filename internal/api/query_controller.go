package api

import (
	"context"
	"net/http"

	"ordercontext/internal/application/query"

	"github.com/labstack/echo/v4"
)

type OrderQueryService interface {
	GetOrders(context.Context) *query.GetOrdersDto
	GetOrder(ctx context.Context, id string) *query.GetOrderDto
}

type OrderQueryController struct {
	orderservice OrderQueryService
}

func NewOrderQueryController(s OrderQueryService) *OrderQueryController {
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
	return handleR(c, http.StatusOK, func(ctx context.Context) (interface{}, error) {
		return o.orderservice.GetOrders(ctx), nil
	})
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
	id := c.Param("id")

	return handleR(c, http.StatusOK, func(ctx context.Context) (interface{}, error) {
		return o.orderservice.GetOrder(ctx, id), nil
	})
}
