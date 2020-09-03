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

func (this *Order) Pay() {
	this.status = Paid
	this.AddEvent(OrderPaidEvent{id: string(this.ID)})
}

func (this *Order) Cancel() {
	this.status = Cancelled
	this.AddEvent(OrderCancelledEvent{id: string(this.ID)})
}

func (this *Order) Ship() error {

	if this.status != Paid {
		return OrderNotPaidError
	}

	this.status = Shipped
	return nil
}
