package command

import (
	"context"
	"orderContext/domain/order"
)

type PayOrderCommand struct {
	OrderId string `validate:"required,min=10"`
}

type PayOrderCommandHandler struct {
	repository order.OrderRepository
}

func NewPayOrderCommandHandler() PayOrderCommandHandler {
	return PayOrderCommandHandler{repository: order.NewOrderRepository()}
}

func (handler PayOrderCommandHandler) Handle(ctx context.Context, cmd PayOrderCommand) error {
	order := handler.repository.Get(cmd.OrderId)
	order.Pay()
	handler.repository.Update(order)
	return nil
}
