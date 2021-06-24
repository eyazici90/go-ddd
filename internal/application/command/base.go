package command

import (
	"context"

	"ordercontext/internal/domain"
)

type (
	GetOrder    func(context.Context, string) (*domain.Order, error)
	GetOrders   func(context.Context) ([]*domain.Order, error)
	CreateOrder func(context.Context, *domain.Order) error
	UpdateOrder func(context.Context, *domain.Order) error

	commandHandlerBase struct {
		getOrder    GetOrder
		updateOrder UpdateOrder
	}
)

func newcommandHandlerBase(getOrder GetOrder, updateOrder UpdateOrder) commandHandlerBase {
	return commandHandlerBase{getOrder, updateOrder}
}

func (handler commandHandlerBase) update(ctx context.Context,
	identifier string,
	when func(*domain.Order)) error {

	o, err := handler.getOrder(ctx, identifier)

	if err != nil {
		return err
	}

	if o == nil {
		return domain.ErrAggregateNotFound
	}
	when(o)

	return handler.updateOrder(ctx, o)
}
