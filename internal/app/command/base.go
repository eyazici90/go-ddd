package command

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/domain"
	"github.com/eyazici90/go-ddd/pkg/aggregate"

	"github.com/pkg/errors"
)

const (
	createCommandKey int = iota
	payCommandKey
	cancelCommandKey
	shipCommandKey
)

type (
	OrderGetter interface {
		Get(context.Context, string) (*domain.Order, error)
	}
	OrderUpdater interface {
		Update(context.Context, *domain.Order) error
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
	fn func(*domain.Order)) error {
	return h.updateErr(ctx, identifier, func(o *domain.Order) error {
		fn(o)
		return nil
	})
}

func (h orderHandler) updateErr(ctx context.Context,
	identifier string,
	fn func(*domain.Order) error) error {
	o, err := h.orderGetter.Get(ctx, identifier)
	if err != nil {
		return errors.Wrap(err, "getting order")
	}
	if o == nil {
		return errors.Wrapf(aggregate.ErrNotFound, "identifier : (%s)", identifier)
	}

	if err := fn(o); err != nil {
		return err
	}

	if err := h.orderUpdater.Update(ctx, o); err != nil {
		return errors.Wrap(err, "updating order")
	}
	return nil
}
