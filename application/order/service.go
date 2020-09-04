package application

import (
	"orderContext/domain/customer"
	"orderContext/domain/order"
	"orderContext/domain/product"

	"github.com/google/uuid"
)

type OrderService interface {
	Create()

	Pay(orderId string)

	Ship(orderId string) error

	Cancel(orderId string)

	GetOrders() []order.Order

	GetOrder(id string) order.Order
}

type service struct {
	repository order.OrderRepository
}

func NewOrderService() OrderService {
	return &service{repository: order.NewOrderRepository()}
}

func (s *service) Create() {
	order := order.NewOrder(order.OrderId(uuid.New().String()), customer.New(), product.New())

	s.repository.Create(*order)
}

func (s *service) Pay(orderId string) {
	order := s.repository.Get(orderId)
	order.Pay()
	s.repository.Update(order)

}

func (s *service) Cancel(orderId string) {
	order := s.repository.Get(orderId)
	order.Pay()
	s.repository.Update(order)
}

func (s *service) Ship(orderId string) error {
	order := s.repository.Get(orderId)
	result := order.Ship()
	if result != nil {
		return result
	}
	s.repository.Update(order)

	return nil
}

func (s *service) GetOrders() []order.Order {
	return s.repository.GetOrders()
}

func (s *service) GetOrder(id string) order.Order {
	return s.repository.Get(id)
}
