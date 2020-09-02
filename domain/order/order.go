package order

import (
	"orderContext/domain/customer"
	"orderContext/domain/product"
	"orderContext/domain/shared"
)

type Status int

const (
	Submitted Status = 1
	Paid      Status = 2
	Shipped   Status = 3
	Cancelled Status = 4
)

type OrderId string

type Order struct {
	shared.AggregateRoot
	ID         OrderId
	customerId customer.CustomerId
	productId  product.ProductId
	status     Status
}

func NewOrder(id OrderId, customerId customer.CustomerId, productId product.ProductId) *Order {
	order := &Order{
		ID:         id,
		customerId: customerId,
		productId:  productId,
	}
	order.AddEvent(OrderCreatedEvent{id: string(id)})
	return order
}

func (o *Order) Pay() {
	o.status = Paid
	o.AddEvent(OrderPaidEvent{id: string(o.ID)})
}

func (o *Order) Cancel() {
	o.status = Cancelled
	o.AddEvent(OrderCancelledEvent{id: string(o.ID)})
}

func (o *Order) Ship() error {

	if o.status != Paid {
		return OrderNotPaidError
	}

	o.status = Shipped
	return nil
}
