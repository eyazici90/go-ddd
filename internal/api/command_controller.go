package api

import (
	"context"
	"net/http"
	"time"

	"ordercontext/internal/application"
	"ordercontext/internal/application/command"
	"ordercontext/internal/application/event"

	"github.com/eyazici90/go-mediator/mediator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OrderCommandController struct {
	sender mediator.Sender
}

func NewOrderCommandController(r application.OrderStore,
	e event.Publisher,
	timeout time.Duration) *OrderCommandController {
	return &OrderCommandController{
		sender: application.NewMediator(r, e, timeout),
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
func (o *OrderCommandController) create(c echo.Context) error {
	return handle(c,
		http.StatusCreated,
		func(ctx context.Context) error {
			return o.sender.Send(ctx, command.CreateOrderCommand{ID: uuid.New().String()})
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
func (o *OrderCommandController) pay(c echo.Context) error {
	id := c.Param("id")

	return handle(c, http.StatusAccepted, func(ctx context.Context) error {
		return o.sender.Send(ctx, command.PayOrderCommand{OrderID: id})
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
func (o *OrderCommandController) cancel(c echo.Context) error {
	id := c.Param("id")

	return handle(c, http.StatusAccepted, func(ctx context.Context) error {
		return o.sender.Send(ctx, command.CancelOrderCommand{OrderID: id})
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
func (o *OrderCommandController) ship(c echo.Context) error {
	id := c.Param("id")

	return handle(c, http.StatusAccepted, func(ctx context.Context) error {
		return o.sender.Send(ctx, command.ShipOrderCommand{OrderID: id})
	})
}
