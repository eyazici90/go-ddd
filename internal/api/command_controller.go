package api

import (
	"context"

	"ordercontext/internal/application"
	"ordercontext/internal/application/command"
	"ordercontext/internal/domain/order"
	"ordercontext/internal/infrastructure"

	"github.com/eyazici90/go-mediator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type orderCommandController struct {
	mediator mediator.Mediator
}

func newOrderCommandController(r order.Repository,
	e infrastructure.EventPublisher,
	timeout int) orderCommandController {
	return orderCommandController{
		mediator: application.NewMediator(r, e, timeout),
	}
}

// CreateOrder godoc
// @Summary Create a order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Router /orders [post]
func (o *orderCommandController) create(c echo.Context) error {
	return create(c, func(ctx context.Context) error {
		return o.mediator.Send(ctx, command.CreateOrderCommand{Id: uuid.New().String()})
	})
}

// PayOrder godoc
// @Summary Pay order
// @Description Pay the order
// @Tags order
// @Accept json
// @Produce json
// @Success 202 {object} string
// @Param id path string true "id"
// @Router /orders/pay/{id} [put]
func (o *orderCommandController) pay(c echo.Context) error {
	return update(c, func(ctx context.Context, id string) error {
		return o.mediator.Send(ctx, command.PayOrderCommand{OrderId: id})
	})
}

// CancelOrder godoc
// @Summary Cancel order
// @Description Cancel the order
// @Tags order
// @Accept json
// @Produce json
// @Success 202 {object} string
// @Param id path string true "id"
// @Router /orders/cancel/{id} [put]
func (o *orderCommandController) cancel(c echo.Context) error {
	return update(c, func(ctx context.Context, id string) error {
		return o.mediator.Send(ctx, command.CancelOrderCommand{OrderId: id})
	})
}

// ShipOrder godoc
// @Summary Ship order
// @Description ship the order
// @Tags order
// @Accept json
// @Produce json
// @Success 202 {object} string
// @Param id path string true "id"
// @Router /orders/ship/{id} [put]
func (o *orderCommandController) ship(c echo.Context) error {
	return update(c, func(ctx context.Context, id string) error {
		return o.mediator.Send(ctx, command.ShipOrderCommand{OrderId: id})
	})
}
