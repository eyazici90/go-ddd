package command

import (
	"context"
	"orderContext/domain/order"
)

type CancelOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type CancelOrderCommandHandler struct {
	commandHandlerBase
}

func (handler CancelOrderCommandHandler) Handle(ctx context.Context, cmd CancelOrderCommand) error {
	return handler.update(ctx, cmd.OrderId, func(order order.Order) {
		order.Cancel()
	})
}
