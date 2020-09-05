package command

import (
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/order"
	"orderContext/domain/product"
)

type CreateOrderCommand struct {
	Id string
}

type CreateOrderCommandHandler struct {
	repository order.OrderRepository
}

func NewCreateOrderCommandHandler() CreateOrderCommandHandler {
	return CreateOrderCommandHandler{repository: order.NewOrderRepository()}
}

func (handler CreateOrderCommandHandler) Handle(cmd CreateOrderCommand) {
	order := order.NewOrder(cmd.Id, customer.New(), product.New(), func() time.Time { return time.Now() })

	handler.repository.Create(*order)
}
