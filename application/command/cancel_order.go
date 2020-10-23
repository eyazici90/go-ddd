package command

import (
	"context"
	"ordercontext/domain/order"
)

type CancelOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type CancelOrderCommandHandler struct {
	commandHandlerBase
}

func (handler CancelOrderCommandHandler) Handle(ctx context.Context, request interface{}) error {
	cmd := request.(CancelOrderCommand)
	return handler.update(ctx, cmd.OrderId, func(order *order.Order) {
		order.Cancel()
	})
}
