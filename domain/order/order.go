package order

import (
	"time"

	"orderContext/core/aggregate"
	"orderContext/domain/customer"
	"orderContext/domain/product"
)

type Status int

const (
	Submitted Status = 1
	Paid      Status = 2
	Shipped   Status = 3
	Cancelled Status = 4
)

type Order struct {
	aggregate.AggregateRoot
	customerId  customer.CustomerId
	productId   product.ProductId
	createdTime time.Time
	status      Status
}

func NewOrder(id string, customerId customer.CustomerId, productId product.ProductId, now aggregate.Now) *Order {
	order := &Order{
		customerId:  customerId,
		productId:   productId,
		createdTime: now(),
	}
	order.ID = id
	order.AddEvent(OrderCreatedEvent{id: id})
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
