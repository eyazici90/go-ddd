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

func (this *service) Create() {
	order := order.NewOrder(order.OrderId(uuid.New().String()), customer.New(), product.New())

	this.repository.Create(*order)
}

func (this *service) Pay(orderId string) {
	order := this.repository.Get(orderId)
	order.Pay()
	this.repository.Update(order)

}

func (this *service) Cancel(orderId string) {
	order := this.repository.Get(orderId)
	order.Pay()
	this.repository.Update(order)
}

func (this *service) Ship(orderId string) error {
	order := this.repository.Get(orderId)
	result := order.Ship()
	if result != nil {
		return result
	}
	this.repository.Update(order)

	return nil
}

func (this *service) GetOrders() []order.Order {
	return this.repository.GetOrders()
}

func (this *service) GetOrder(id string) order.Order {
	return this.repository.Get(id)
}
