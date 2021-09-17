package command

import (
	"context"
	"time"

	"ordercontext/internal/domain/order"
	"ordercontext/pkg/aggregate"

	"github.com/eyazici90/go-mediator/mediator"
	"github.com/pkg/errors"
)

type CreateOrderFn func(context.Context, *order.Order) error

type CreateOrder struct {
	ID string `validate:"required,min=10"`
}

func (CreateOrder) Key() string { return "CreateOrder" }

type CreateOrderHandler struct {
	createOrderFn CreateOrderFn
}

func NewCreateOrderHandler(createOrderFn CreateOrderFn) CreateOrderHandler {
	return CreateOrderHandler{createOrderFn}
}

func (h CreateOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CreateOrder)
	if err := checkType(ok); err != nil {
		return err
	}

	ordr, err := order.NewOrder(order.ID(cmd.ID), order.NewCustomerID(), order.NewProductID(), time.Now,
		order.Submitted, aggregate.NewVersion())

	if err != nil {
		return errors.Wrap(err, "create order handle failed")
	}

	return h.createOrderFn(ctx, ordr)
}
