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
	repository order.OrderRepository
}

func NewCreateOrderCommandHandler(r order.OrderRepository) CreateOrderCommandHandler {
	return CreateOrderCommandHandler{repository: r}
}

func (handler CreateOrderCommandHandler) Handle(_ context.Context, cmd CreateOrderCommand) error {
	order := order.NewOrder(cmd.Id, customer.New(), product.New(), func() time.Time { return time.Now() })

	handler.repository.Create(*order)

	return nil
}
