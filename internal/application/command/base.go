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
	return handler.updateErr(ctx, identifier, func(o *domain.Order) error {
		fn(o)
		return nil
	})
}

func (handler commandHandler) updateErr(ctx context.Context,
	identifier string,
	fn func(*domain.Order) error) error {
	order, err := handler.getOrder(ctx, identifier)

	if err != nil {
		return errors.Wrap(err, "get order failed")
	}
	if order == nil {
		return errors.Wrapf(domain.ErrAggregateNotFound, "identifier : (%s)", identifier)
	}

	if err := fn(order); err != nil {
		return err
	}
	return handler.updateOrder(ctx, order)
}
