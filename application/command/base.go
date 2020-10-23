package command

import (
	"context"
	"ordercontext/domain/order"
)

type (
	GetOrder    func(context.Context, string) (*order.Order, error)
	GetOrders   func(context.Context) ([]*order.Order, error)
	CreateOrder func(context.Context, *order.Order) error
	UpdateOrder func(context.Context, *order.Order) error

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
	when func(*order.Order)) error {

	existingOrder, err := handler.getOrder(ctx, identifier)

	if err != nil {
		return err
	}

	if existingOrder == nil {
		return order.ErrAggregateNotFound
	}
	when(existingOrder)

	handler.updateOrder(ctx, existingOrder)

	return nil
}
