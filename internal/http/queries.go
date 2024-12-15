package http

import (
	"context"
	"net/http"

	"github.com/eyazici90/go-ddd/internal/app/query"
	"github.com/labstack/echo/v4"
)

type OrderQuerier interface {
	GetOrders(ctx context.Context) *query.GetOrdersDto
	GetOrder(ctx context.Context, id string) *query.GetOrderDto
}

type QueryController struct {
	orderQuerier OrderQuerier
}

func NewQueryController(s OrderQuerier) *QueryController {
	return &QueryController{
		orderQuerier: s,
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
func (oqc *QueryController) getOrders(c echo.Context) error {
	return handleR(c, http.StatusOK, func(ctx context.Context) (interface{}, error) {
		return oqc.orderQuerier.GetOrders(ctx), nil
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
func (oqc *QueryController) getOrder(c echo.Context) error {
	id := c.Param("id")

	return handleR(c, http.StatusOK, func(ctx context.Context) (interface{}, error) {
		return oqc.orderQuerier.GetOrder(ctx, id), nil
	})
}
