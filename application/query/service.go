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

func NewOrderQueryService(r order.OrderRepository) OrderQueryService {
	return &service{repository: r}
}

func (s *service) GetOrders() []order.Order {
	return s.repository.GetOrders()
}

func (s *service) GetOrder(id string) order.Order {
	return s.repository.Get(id)
}
