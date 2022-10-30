package query

import (
	"context"
	"net/http"

	"github.com/eyazici90/go-ddd/internal/app/feature"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetOrders(ctx context.Context) *GetOrdersDto
	GetOrder(ctx context.Context, id string) *GetOrderDto
}

type OrderController struct {
	os Service
}

func NewOrderController(s Service) OrderController {
	return OrderController{
		os: s,
	}
}

// GetOrders godoc
// @Summary Get orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} query.GetOrdersDto
// @Router /orders [get]
func (o OrderController) GetOrders(c echo.Context) error {
	return feature.HandleR(c, http.StatusOK, func(ctx context.Context) (interface{}, error) {
		return o.os.GetOrders(ctx), nil
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
func (o OrderController) GetOrder(c echo.Context) error {
	id := c.Param("id")

	return feature.HandleR(c, http.StatusOK, func(ctx context.Context) (interface{}, error) {
		return o.os.GetOrder(ctx, id), nil
	})
}
