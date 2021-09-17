package command

import (
	"context"
	"time"

	"ordercontext/internal/domain/order"
	"ordercontext/pkg/aggregate"

	"github.com/eyazici90/go-mediator/mediator"
	"github.com/pkg/errors"
)

type CreateOrder func(context.Context, *order.Order) error

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

func (h CreateOrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CreateOrderCommand)
	if err := checkType(ok); err != nil {
		return err
	}

	ordr, err := order.NewOrder(order.ID(cmd.ID), order.NewCustomerID(), order.NewProductID(), func() time.Time { return time.Now() },
		order.Submitted, aggregate.NewVersion())

	if err != nil {
		return errors.Wrap(err, "create order handle failed")
	}

	return h.createOrder(ctx, ordr)
}
