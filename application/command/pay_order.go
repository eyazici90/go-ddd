package command

import (
	"orderContext/domain/order"
)

type PayOrderCommand struct {
	OrderId string
}

type PayOrderCommandHandler struct {
	repository order.OrderRepository
}

func (handler PayOrderCommandHandler) Handle(cmd PayOrderCommand) {
	order := handler.repository.Get(cmd.OrderId)
	order.Pay()
	handler.repository.Update(order)
}
