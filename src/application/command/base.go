package command

import (
	"context"
	"orderContext/domain/order"
)

type commandHandlerBase struct {
	repository order.OrderRepository
}

func newcommandHandlerBase(r order.OrderRepository) commandHandlerBase {
	return commandHandlerBase{
		repository: r,
	}
}

func (handler commandHandlerBase) update(ctx context.Context,
	identifier string,
	when func(order.Order)) error {

	existingOrder := handler.repository.Get(ctx, identifier)

	if existingOrder == nil {
		return order.AggregateNotFound
	}
	when(existingOrder)

	handler.repository.Update(ctx, existingOrder)

	return nil
}
