package command

import (
	"context"
	"fmt"

	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
)

const (
	createCommandKey int = iota + 1
	payCommandKey
	cancelCommandKey
	shipCommandKey
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

func newOrderHandler(getter OrderGetter, updater OrderUpdater) orderHandler {
	return orderHandler{getter, updater}
}

func (h orderHandler) update(ctx context.Context,
	id string,
	fn func(*order.Order)) error {
	return h.updateErr(ctx, id, func(o *order.Order) error {
		fn(o)
		return nil
	})
}

func (h orderHandler) updateErr(ctx context.Context,
	id string,
	fn func(*order.Order) error) error {
	o, err := h.orderGetter.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("getting order: %w", err)
	}
	if o == nil {
		return fmt.Errorf("id: (%s): %w", id, aggregate.ErrNotFound)
	}
	if err = fn(o); err != nil {
		return err
	}
	if err = h.orderUpdater.Update(ctx, o); err != nil {
		return fmt.Errorf("updating order: %w", err)
	}
	return nil
}
