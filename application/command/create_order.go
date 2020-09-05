package command

import (
	"time"

	"github.com/google/uuid"

	"orderContext/domain/customer"
	"orderContext/domain/order"
	"orderContext/domain/product"
)

type CreateOrderCommand struct {
}

type CreateOrderCommandHandler struct {
	repository order.OrderRepository
}

type CreateOrderCommandValidator struct {
	cmd CreateOrderCommand
}

func NewCreateOrderCommandHandler() CreateOrderCommandHandler {
	return CreateOrderCommandHandler{repository: order.NewOrderRepository()}
}

func (handler CreateOrderCommandHandler) Handle(cmd CreateOrderCommand) {
	order := order.NewOrder(uuid.New().String(), customer.New(), product.New(), func() time.Time { return time.Now() })

	handler.repository.Create(*order)
}

func (validator *CreateOrderCommandValidator) Validate(cmd CreateOrderCommand) error {
	return nil
}
