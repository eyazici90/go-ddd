package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/app/event"
	"github.com/eyazici90/go-ddd/internal/domain"

	"github.com/eyazici90/go-mediator/pkg/mediator"
)

type ShipOrder struct {
	OrderID string `validate:"required,min=10"`
}

func (ShipOrder) Key() string { return "ShipOrder" }

type ShipOrderHandler struct {
	orderHandler
	eventPublisher event.Publisher
}

func NewShipOrderHandler(orderGetter OrderGetter,
	orderUpdater OrderUpdater,
	e event.Publisher) ShipOrderHandler {
	return ShipOrderHandler{
		orderHandler:   newOrderHandler(orderGetter, orderUpdater),
		eventPublisher: e,
	}
}

func (h ShipOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(ShipOrder)
	if err := checkType(ok); err != nil {
		return err
	}

	var ord *domain.Order
	if err := h.updateErr(ctx, cmd.OrderID, func(o *domain.Order) error {
		ord = o
		return ord.Ship()
	}); err != nil {
		return err
	}

	h.eventPublisher.PublishAll(ord.Events())

	return nil
}
