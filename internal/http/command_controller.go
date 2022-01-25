package http

import (
	"context"
	"net/http"
	"time"

	"github.com/eyazici90/go-ddd/internal/app"
	"github.com/eyazici90/go-ddd/internal/app/command"
	"github.com/eyazici90/go-ddd/internal/app/event"

	"github.com/eyazici90/go-mediator/mediator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CommandController struct {
	sender mediator.Sender
}

func NewCommandController(r app.OrderStore,
	e event.Publisher,
	timeout time.Duration) (CommandController, error) {
	m, err := app.NewMediator(r, e, timeout)
	if err != nil {
		return CommandController{}, err
	}
	return CommandController{
		sender: m,
	}, nil
}

// CreateOrder godoc
// @Summary Create an order
// @Description Create a new order
// @Tags order
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Router /orders [post]
func (o CommandController) create(c echo.Context) error {
	return handle(c,
		http.StatusCreated,
		func(ctx context.Context) error {
			return o.sender.Send(ctx, command.CreateOrder{ID: uuid.New().String()})
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
func (o CommandController) pay(c echo.Context) error {
	id := c.Param("id")

	return handle(c, http.StatusAccepted, func(ctx context.Context) error {
		return o.sender.Send(ctx, command.PayOrder{OrderID: id})
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
func (o CommandController) cancel(c echo.Context) error {
	id := c.Param("id")

	return handle(c, http.StatusAccepted, func(ctx context.Context) error {
		return o.sender.Send(ctx, command.CancelOrder{OrderID: id})
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
func (o CommandController) ship(c echo.Context) error {
	id := c.Param("id")

	return handle(c, http.StatusAccepted, func(ctx context.Context) error {
		return o.sender.Send(ctx, command.ShipOrder{OrderID: id})
	})
}
