package api

import (
	"context"
	"orderContext/application/behaviour"
	"orderContext/domain/order"
	"orderContext/infrastructure"

	"orderContext/application/command"
	"orderContext/core/mediator"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type orderCommandController struct {
	mediator mediator.Mediator
}

func newOrderCommandController(r order.OrderRepository, e infrastructure.EventPublisher) orderCommandController {
	m := mediator.NewMediator().
		UseBehaviour(behaviour.NewCancellator()).
		UseBehaviour(behaviour.NewLogger()).
		UseBehaviour(behaviour.NewValidator()).
		RegisterHandler(command.NewCreateOrderCommandHandler(r)).
		RegisterHandler(command.NewPayOrderCommandHandler(r)).
		RegisterHandler(command.NewShipOrderCommandHandler(r, e))

	return orderCommandController{
		mediator: m,
	}
}

// CreateOrder godoc
// @Summary Create a order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Router /order [post]
func (o *orderCommandController) create(c echo.Context) error {
	return create(c, func(ctx context.Context) {
		o.mediator.Send(ctx, command.CreateOrderCommand{Id: uuid.New().String()})
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
// @Router /order/pay/{id} [put]
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
// @Router /order/cancel/{id} [put]
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
// @Router /order/ship/{id} [put]
func (o *orderCommandController) ship(c echo.Context) error {
	return update(c, func(ctx context.Context, id string) error {
		return o.mediator.Send(ctx, command.ShipOrderCommand{OrderId: id})
	})
}
