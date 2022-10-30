package create

import (
	"context"
	"fmt"
	"time"

	"github.com/eyazici90/go-ddd/internal/app/feature"
	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/eyazici90/go-mediator/mediator"
)

const commandKey int = 1

type OrderCreator interface {
	Create(context.Context, *order.Order) error
}

type OrderCommand struct {
	OrderID string `validate:"required,min=10"`
}

func (OrderCommand) Key() int { return commandKey }

type OrderCommandHandler struct {
	orderCreator OrderCreator
}

func NewOrderCommandHandler(creator OrderCreator) OrderCommandHandler {
	return OrderCommandHandler{orderCreator: creator}
}

func (h OrderCommandHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(OrderCommand)
	if !ok {
		return feature.ErrInvalidCommand
	}

	ord, err := order.New(order.ID(cmd.OrderID), order.NewCustomerID(), order.NewProductID(), time.Now,
		order.Submitted, aggregate.NewVersion())
	if err != nil {
		return fmt.Errorf("new order: %w", err)
	}

	return h.orderCreator.Create(ctx, ord)
}
