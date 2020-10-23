package command

import (
	"context"
	"ordercontext/shared/aggregate"
	"time"

	"ordercontext/domain/customer"
	"ordercontext/domain/order"
	"ordercontext/domain/product"
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

func (handler CreateOrderCommandHandler) Handle(ctx context.Context, request interface{}) error {
	cmd := request.(CreateOrderCommand)
	order, err := order.NewOrder(order.OrderID(cmd.Id), customer.New(), product.New(), func() time.Time { return time.Now() },
		order.Submitted, aggregate.NewVersion())

	if err != nil {
		return err
	}

	handler.createOrder(ctx, order)

	return nil
}
