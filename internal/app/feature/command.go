package feature

import (
	"context"
	"fmt"

	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
)

type (
	OrderGetter interface {
		Get(context.Context, string) (*order.Order, error)
	}
	OrderUpdater interface {
		Update(context.Context, *order.Order) error
	}
)

type OrderHandler struct {
	OrderGetter  OrderGetter
	OrderUpdater OrderUpdater
}

func (h OrderHandler) Update(ctx context.Context,
	id string,
	fn func(*order.Order),
) error {
	return h.UpdateErr(ctx, id, func(o *order.Order) error {
		fn(o)
		return nil
	})
}

func (h OrderHandler) UpdateErr(ctx context.Context,
	id string,
	fn func(*order.Order) error,
) error {
	o, err := h.OrderGetter.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("getting order: %w", err)
	}
	if o == nil {
		return fmt.Errorf("id: (%s): %w", id, aggregate.ErrNotFound)
	}
	if err = fn(o); err != nil {
		return err
	}
	if err = h.OrderUpdater.Update(ctx, o); err != nil {
		return fmt.Errorf("updating order: %w", err)
	}
	return nil
}
