package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/app/event"
	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-mediator/mediator"
)

type ShipOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (ShipOrder) Key() int { return shipCommandKey }

type ShipOrderHandler struct {
	orderHandler
	eventPublisher event.Publisher
}

func NewShipOrderHandler(getter OrderGetter,
	updater OrderUpdater,
	e event.Publisher) ShipOrderHandler {
	return ShipOrderHandler{
		orderHandler:   newOrderHandler(getter, updater),
		eventPublisher: e,
	}
}

func (h ShipOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(ShipOrder)
	if !ok {
		return ErrInvalidCommand
	}

	var ord *order.Order
	if err := h.updateErr(ctx, cmd.OrderID, func(o *order.Order) error {
		ord = o
		return ord.Ship()
	}); err != nil {
		return err
	}

	h.eventPublisher.PublishAll(ord.Events())

	return nil
}
