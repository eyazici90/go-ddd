package command

import (
	"orderContext/domain/order"
)

type PayOrderCommand struct {
	OrderId string `validate:"required,min=3"`
}

type PayOrderCommandHandler struct {
	repository order.OrderRepository
}

func NewPayOrderCommandHandler() PayOrderCommandHandler {
	return PayOrderCommandHandler{repository: order.NewOrderRepository()}
}

func (handler PayOrderCommandHandler) Handle(cmd PayOrderCommand) {
	order := handler.repository.Get(cmd.OrderId)
	order.Pay()
	handler.repository.Update(order)
}
