package command

import (
	"context"

	"ordercontext/internal/domain/order"
	"ordercontext/internal/infrastructure"

	"github.com/eyazici90/go-mediator/mediator"
)

type ShipOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (ShipOrderCommand) Key() string { return "ShipOrderCommand" }

type ShipOrderCommandHandler struct {
	repository     order.Repository
	eventPublisher infrastructure.EventPublisher
}

func NewShipOrderCommandHandler(r order.Repository, e infrastructure.EventPublisher) ShipOrderCommandHandler {
	return ShipOrderCommandHandler{
		repository:     r,
		eventPublisher: e,
	}
}

func (h ShipOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd := msg.(ShipOrderCommand)
	o, err := h.repository.Get(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	err = o.Ship()

	if err != nil {
		return err
	}

	h.repository.Update(ctx, o)

	h.eventPublisher.PublishAll(o.Events())

	return nil
}
