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
	commandHandlerBase
}

func (h CancelOrderCommandHandler) Handle(ctx context.Context, req mediator.Message) error {
	cmd := req.(CancelOrderCommand)
	return h.update(ctx, cmd.OrderID, func(o *order.Order) {
		o.Cancel()
	})
}
