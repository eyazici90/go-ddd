package command

import (
	"context"
	"orderContext/domain/order"
)

type PayOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type PayOrderCommandHandler struct {
	commandHandlerBase
}

func NewPayOrderCommandHandler(r order.Repository) PayOrderCommandHandler {
	return PayOrderCommandHandler{
		commandHandlerBase: newcommandHandlerBase(r),
	}
}

func (handler PayOrderCommandHandler) Handle(ctx context.Context, cmd PayOrderCommand) error {
	return handler.update(ctx, cmd.OrderId, func(order order.Order) {
		order.Pay()
	})
}
