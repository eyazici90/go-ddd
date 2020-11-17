package command

import (
	"context"

	"ordercontext/internal/domain/order"

	"github.com/eyazici90/go-mediator"
)

type CancelOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

func (CancelOrderCommand) Key() string { return "CancelOrderCommand " }

type CancelOrderCommandHandler struct {
	commandHandlerBase
}

func (handler CancelOrderCommandHandler) Handle(ctx context.Context, request mediator.Message) error {
	cmd := request.(CancelOrderCommand)
	return handler.update(ctx, cmd.OrderId, func(order *order.Order) {
		order.Cancel()
	})
}
