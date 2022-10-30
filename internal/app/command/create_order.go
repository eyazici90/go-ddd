package command

import (
	"context"
	"fmt"
	"time"

	"github.com/eyazici90/go-ddd/internal/domain"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/eyazici90/go-mediator/mediator"
)

type OrderCreator interface {
	Create(context.Context, *domain.Order) error
}

type CreateOrder struct {
	ID string `validate:"required,min=10"`
}

func (CreateOrder) Key() int { return createCommandKey }

type CreateOrderHandler struct {
	creator OrderCreator
}

func NewCreateOrderHandler(creator OrderCreator) CreateOrderHandler {
	return CreateOrderHandler{creator: creator}
}

func (h CreateOrderHandler) Handle(ctx context.Context, msg mediator.Message) error {
	cmd, ok := msg.(CreateOrder)
	if !ok {
		return ErrInvalidCommand
	}

	order, err := domain.NewOrder(domain.OrderID(cmd.ID), domain.NewCustomerID(), domain.NewProductID(), time.Now,
		domain.Submitted, aggregate.NewVersion())
	if err != nil {
		return fmt.Errorf("new order: %w", err)
	}

	return h.creator.Create(ctx, order)
}
