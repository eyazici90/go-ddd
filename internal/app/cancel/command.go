package cancel

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/app/feature"
	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-mediator/mediator"
)

const commandKey int = 3

type OrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (OrderCommand) Key() int { return commandKey }

type OrderCommandHandler struct {
	feature.OrderHandler
}

func NewOrderCommandHandler(getter feature.OrderGetter,
	updater feature.OrderUpdater,
) OrderCommandHandler {
	return OrderCommandHandler{
		OrderHandler: feature.OrderHandler{
			OrderGetter:  getter,
			OrderUpdater: updater,
		},
	}
}

func (h OrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(OrderCommand)
	if !ok {
		return feature.ErrInvalidCommand
	}
	return h.Update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Cancel()
	})
}
