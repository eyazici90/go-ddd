package command

import (
	"context"
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/order"
	"orderContext/domain/product"
)

type CreateOrderCommand struct {
	Id string `validate:"required,min=10"`
}

type CreateOrderCommandHandler struct {
	createOrder CreateOrder
}

func NewCreateOrderCommandHandler(createOrder CreateOrder) CreateOrderCommandHandler {
	return CreateOrderCommandHandler{createOrder}
}

func (handler CreateOrderCommandHandler) Handle(ctx context.Context, cmd CreateOrderCommand) error {
	order, err := order.NewOrder(order.OrderId(cmd.Id), customer.New(), product.New(), func() time.Time { return time.Now() })

	if err != nil {
		return err
	}

	handler.createOrder(ctx, order)

	return nil
}
