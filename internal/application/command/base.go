package command

import (
	"context"

	"ordercontext/internal/domain/order"
	"ordercontext/pkg/aggregate"

	"github.com/pkg/errors"
)

type (
	GetOrder    func(context.Context, string) (*order.Order, error)
	UpdateOrder func(context.Context, *order.Order) error

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
	fn func(*order.Order)) error {
	return handler.updateErr(ctx, identifier, func(o *order.Order) error {
		fn(o)
		return nil
	})
}

func (handler commandHandler) updateErr(ctx context.Context,
	identifier string,
	fn func(*order.Order) error) error {
	o, err := handler.getOrder(ctx, identifier)

	if err != nil {
		return errors.Wrap(err, "get order failed")
	}
	if o == nil {
		return errors.Wrapf(aggregate.ErrNotFound, "identifier : (%s)", identifier)
	}

	if err := fn(o); err != nil {
		return err
	}
	return handler.updateOrder(ctx, o)
}
