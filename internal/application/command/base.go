package command

import (
	"context"

	"ordercontext/internal/domain"

	"github.com/pkg/errors"
)

type (
	GetOrder    func(context.Context, string) (*domain.Order, error)
	UpdateOrder func(context.Context, *domain.Order) error

	commandHandler struct {
		getOrder    GetOrder
		updateOrder UpdateOrder
	}
)

func newcommandHandlerBase(getOrder GetOrder, updateOrder UpdateOrder) commandHandler {
	return commandHandler{getOrder, updateOrder}
}

func (handler commandHandler) update(ctx context.Context,
	identifier string,
	fn func(*domain.Order)) error {
	order, err := handler.getOrder(ctx, identifier)

	if err != nil {
		return errors.Wrap(err, "get order failed")
	}
	if order == nil {
		return errors.Wrapf(domain.ErrAggregateNotFound, "identifier : (%s)", identifier)
	}

	fn(order)
	return handler.updateOrder(ctx, order)
}
