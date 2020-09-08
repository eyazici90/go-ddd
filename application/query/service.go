package query

import (
	"orderContext/domain/order"
)

type OrderQueryService interface {
	GetOrders() []order.Order

	GetOrder(id string) order.Order
}

type service struct {
	repository order.OrderRepository
}

func NewOrderQueryService() OrderQueryService {
	return &service{repository: order.NewOrderRepository()}
}

func (s *service) GetOrders() []order.Order {
	return s.repository.GetOrders()
}

func (s *service) GetOrder(id string) order.Order {
	return s.repository.Get(id)
}
