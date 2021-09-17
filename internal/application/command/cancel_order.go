package command

import (
	"context"
	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator/mediator"
)

type CancelOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (CancelOrderCommand) Key() string { return "CancelOrderCommand" }

type CancelOrderCommandHandler struct {
	commandHandler
}

func (h CancelOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CancelOrderCommand)
	if err := checkType(ok); err != nil {
		return err
	}
	return h.update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Cancel()
	})
}
