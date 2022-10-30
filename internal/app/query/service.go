package query

import (
	"context"

	"github.com/eyazici90/go-ddd/internal/order"
)

type OrderQueryStore interface {
	GetAll(context.Context) ([]*order.Order, error)
	Get(ctx context.Context, id string) (*order.Order, error)
}

type service struct {
	store OrderQueryStore
}

func NewOrderService(store OrderQueryStore) *service {
	return &service{store}
}

func (s *service) GetOrders(ctx context.Context) *GetOrdersDto {
	orders, err := s.store.GetAll(ctx)
	if err != nil {
		return nil
	}
	oViews := mapToAll(orders)

	result := &GetOrdersDto{oViews}

	return result
}

func (s *service) GetOrder(ctx context.Context, id string) *GetOrderDto {
	ord, err := s.store.Get(ctx, id)
	if err != nil {
		return nil
	}
	oView := mapTo(ord)

	result := &GetOrderDto{oView}
	return result
}
