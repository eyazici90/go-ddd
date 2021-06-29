package command

import (
	"context"

	"ordercontext/internal/application/event"
	"ordercontext/internal/domain"

	"github.com/eyazici90/go-mediator/mediator"
	"github.com/pkg/errors"
)

type ShipOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (ShipOrderCommand) Key() string { return "ShipOrderCommand" }

type ShipOrderCommandHandler struct {
	repository     domain.OrderRepository
	eventPublisher event.Publisher
}

func NewShipOrderCommandHandler(r domain.OrderRepository, e event.Publisher) ShipOrderCommandHandler {
	return ShipOrderCommandHandler{
		repository:     r,
		eventPublisher: e,
	}
}

func (h ShipOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(ShipOrderCommand)
	if err := checkType(ok); err != nil {
		return err
	}
	order, err := h.repository.Get(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	err = order.Ship()

	if err != nil {
		return errors.Wrap(err, "ship handle failed")
	}

	if err := h.repository.Update(ctx, order); err != nil {
		return err
	}

	h.eventPublisher.PublishAll(order.Events())

	return nil
}
