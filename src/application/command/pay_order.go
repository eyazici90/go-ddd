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

func NewPayOrderCommandHandler(getOrder GetOrder, updateOrder UpdateOrder) PayOrderCommandHandler {
	return PayOrderCommandHandler{
		commandHandlerBase: newcommandHandlerBase(getOrder, updateOrder),
	}
}

func (handler PayOrderCommandHandler) Handle(ctx context.Context, request interface{}) error {
	cmd := request.(PayOrderCommand)
	return handler.update(ctx, cmd.OrderId, func(order order.Order) {
		order.Pay()
	})
}
