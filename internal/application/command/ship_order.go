package command

import (
	"context"

	"ordercontext/internal/application/event"
	"ordercontext/internal/domain"

	"github.com/eyazici90/go-mediator/mediator"
)

type ShipOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (ShipOrderCommand) Key() string { return "ShipOrderCommand" }

type ShipOrderCommandHandler struct {
	commandHandler
	eventPublisher event.Publisher
}

func NewShipOrderCommandHandler(getOrder GetOrder,
	updateOrder UpdateOrder,
	e event.Publisher) ShipOrderCommandHandler {
	return ShipOrderCommandHandler{
		commandHandler: newcommandHandlerBase(getOrder, updateOrder),
		eventPublisher: e,
	}
}

func (h ShipOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(ShipOrderCommand)
	if err := checkType(ok); err != nil {
		return err
	}

	var order *domain.Order
	if err := h.updateErr(ctx, cmd.OrderID, func(o *domain.Order) error {
		order = o
		return order.Ship()
	}); err != nil {
		return err
	}

	h.eventPublisher.PublishAll(order.Events())

	return nil
}
