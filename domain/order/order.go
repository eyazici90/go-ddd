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

func NewOrder(id string, customerId customer.CustomerId, productId product.ProductId, now aggregate.Now) (*Order, error) {
	order := &Order{
		customerId:  customerId,
		productId:   productId,
		createdTime: now(),
	}
	order.ID = id

	err := ValidateState(order)

	if err != nil {
		return nil, err
	}

	order.AddEvent(OrderCreatedEvent{id: id})

	return order, nil
}

func ValidateState(o *Order) error {
	if o.ID == "" || o.customerId == "" {
		return InvalidValueError
	}
	return nil
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

func (o *Order) Status() Status { return o.status }
