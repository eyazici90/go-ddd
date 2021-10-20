package command

import (
	"context"

	"ordercontext/internal/domain/order"
	"ordercontext/pkg/aggregate"

	"github.com/pkg/errors"
)

type (
	OrderGetter interface {
		Get(context.Context, string) (*order.Order, error)
	}
	OrderUpdater interface {
		Update(context.Context, *order.Order) error
	}
)

type orderHandler struct {
	orderGetter  OrderGetter
	orderUpdater OrderUpdater
}

func newOrderHandler(orderGetter OrderGetter, orderUpdater OrderUpdater) orderHandler {
	return orderHandler{orderGetter, orderUpdater}
}

func (h orderHandler) update(ctx context.Context,
	identifier string,
	fn func(*order.Order)) error {
	return h.updateErr(ctx, identifier, func(o *order.Order) error {
		fn(o)
		return nil
	})
}

func (h orderHandler) updateErr(ctx context.Context,
	identifier string,
	fn func(*order.Order) error) error {
	o, err := h.orderGetter.Get(ctx, identifier)

	if err != nil {
		return errors.Wrap(err, "get order failed")
	}
	if o == nil {
		return errors.Wrapf(aggregate.ErrNotFound, "identifier : (%s)", identifier)
	}

	if err := fn(o); err != nil {
		return err
	}
	return h.orderUpdater.Update(ctx, o)
}
