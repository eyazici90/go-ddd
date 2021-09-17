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

	orderHandler struct {
		getOrder    GetOrder
		updateOrder UpdateOrder
	}
)

func newOrderHandler(getOrder GetOrder, updateOrder UpdateOrder) orderHandler {
	return orderHandler{getOrder, updateOrder}
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
	o, err := h.getOrder(ctx, identifier)

	if err != nil {
		return errors.Wrap(err, "get order failed")
	}
	if o == nil {
		return errors.Wrapf(aggregate.ErrNotFound, "identifier : (%s)", identifier)
	}

	if err := fn(o); err != nil {
		return err
	}
	return h.updateOrder(ctx, o)
}
