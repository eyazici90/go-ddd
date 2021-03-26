package command

import (
	"context"
	"ordercontext/internal/domain"

	"github.com/eyazici90/go-mediator/mediator"
)

type CancelOrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (CancelOrderCommand) Key() string { return "CancelOrderCommand" }

type CancelOrderCommandHandler struct {
	commandHandlerBase
}

func (h CancelOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd := msg.(CancelOrderCommand)
	return h.update(ctx, cmd.OrderID, func(o *domain.Order) {
		o.Cancel()
	})
}
