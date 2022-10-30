package ship

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/app/feature"
	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-mediator/mediator"
)

const commandKey int = 2

type OrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (OrderCommand) Key() int { return commandKey }

type OrderCommandHandler struct {
	feature.OrderHandler
	eventPublisher feature.EventPublisher
}

func NewOrderCommandHandler(getter feature.OrderGetter,
	updater feature.OrderUpdater,
	e feature.EventPublisher,
) OrderCommandHandler {
	return OrderCommandHandler{
		OrderHandler: feature.OrderHandler{
			OrderGetter:  getter,
			OrderUpdater: updater,
		},
		eventPublisher: e,
	}
}

func (h OrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(OrderCommand)
	if !ok {
		return feature.ErrInvalidCommand
	}

	var ord *order.Order
	if err := h.UpdateErr(ctx, cmd.OrderID, func(o *order.Order) error {
		ord = o
		return ord.Ship()
	}); err != nil {
		return err
	}

	h.eventPublisher.PublishAll(ord.Events())

	return nil
}
