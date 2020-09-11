package query

import (
	"context"
	"orderContext/domain/order"
)

type OrderQueryService interface {
	GetOrders(context.Context) []order.Order

	GetOrder(ctx context.Context, id string) order.Order
}

type service struct {
	repository order.OrderRepository
}

func NewOrderQueryService(r order.OrderRepository) OrderQueryService {
	return &service{repository: r}
}

func (s *service) GetOrders(ctx context.Context) []order.Order {
	return s.repository.GetOrders(ctx)
}

func (s *service) GetOrder(ctx context.Context, id string) order.Order {
	return s.repository.Get(ctx, id)
}
