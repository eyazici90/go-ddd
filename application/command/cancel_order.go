package command

import (
	"orderContext/domain/order"
)

type CancelOrderCommand struct {
	OrderId string
}

type CancelOrderCommandHandler struct {
	repository order.OrderRepository
}

func (handler CancelOrderCommandHandler) Handle(cmd CancelOrderCommand) {
	order := handler.repository.Get(cmd.OrderId)
	order.Cancel()
	handler.repository.Update(order)
}
