package command

import (
	"context"
	"time"

	"ordercontext/internal/domain/customer"
	"ordercontext/internal/domain/order"
	"ordercontext/internal/domain/product"
	"ordercontext/pkg/aggregate"

	"github.com/eyazici90/go-mediator"
)

type CreateOrderCommand struct {
	Id string `validate:"required,min=10"`
}

func (CreateOrderCommand) Key() string { return "CreateOrderCommand " }

type CreateOrderCommandHandler struct {
	createOrder CreateOrder
}

func NewCreateOrderCommandHandler(createOrder CreateOrder) CreateOrderCommandHandler {
	return CreateOrderCommandHandler{createOrder}
}

func (handler CreateOrderCommandHandler) Handle(ctx context.Context, request mediator.Message) error {
	cmd := request.(CreateOrderCommand)
	order, err := order.NewOrder(order.OrderID(cmd.Id), customer.New(), product.New(), func() time.Time { return time.Now() },
		order.Submitted, aggregate.NewVersion())

	if err != nil {
		return err
	}

	handler.createOrder(ctx, order)

	return nil
}
