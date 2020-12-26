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
	ID string `validate:"required,min=10"`
}

func (CreateOrderCommand) Key() string { return "CreateOrderCommand " }

type CreateOrderCommandHandler struct {
	createOrder CreateOrder
}

func NewCreateOrderCommandHandler(createOrder CreateOrder) CreateOrderCommandHandler {
	return CreateOrderCommandHandler{createOrder}
}

func (h CreateOrderCommandHandler) Handle(ctx context.Context, req mediator.Message) error {
	cmd := req.(CreateOrderCommand)
	o, err := order.NewOrder(order.ID(cmd.ID), customer.New(), product.New(), func() time.Time { return time.Now() },
		order.Submitted, aggregate.NewVersion())

	if err != nil {
		return err
	}

	h.createOrder(ctx, o)

	return nil
}
