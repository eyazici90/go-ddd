package command

import (
	"context"
	"orderContext/domain/order"
)

type (
	GetOrder    func(context.Context, string) *order.Order
	GetOrders   func(context.Context) []*order.Order
	CreateOrder func(context.Context, *order.Order)
	UpdateOrder func(context.Context, *order.Order)

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

	existingOrder := handler.getOrder(ctx, identifier)

	if existingOrder == nil {
		return order.AggregateNotFound
	}
	when(existingOrder)

	handler.updateOrder(ctx, existingOrder)

	return nil
}
