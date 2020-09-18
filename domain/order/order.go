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

type Order interface {
	aggregate.EventTracker
	Pay()
	Cancel()
	Ship() error
	Status() Status
	Id() string
}

type order struct {
	aggregate.AggregateRoot
	customerId  customer.CustomerId
	productId   product.ProductId
	createdTime time.Time
	status      Status
}

func NewOrder(id string, customerId customer.CustomerId, productId product.ProductId, now aggregate.Now) (Order, error) {
	o := &order{
		customerId:  customerId,
		productId:   productId,
		createdTime: now(),
	}
	o.ID = id

	err := ValidateState(o)

	if err != nil {
		return nil, err
	}

	o.AddEvent(OrderCreatedEvent{id: id})

	return o, nil
}

func ValidateState(o *order) error {
	if o.ID == "" || o.customerId == "" {
		return InvalidValueError
	}
	return nil
}

func (o *order) Pay() {
	o.status = Paid
	o.AddEvent(OrderPaidEvent{id: string(o.ID)})
}

func (o *order) Cancel() {
	o.status = Cancelled
	o.AddEvent(OrderCancelledEvent{id: string(o.ID)})
}

func (o *order) Ship() error {

	if o.status != Paid {
		return OrderNotPaidError
	}

	o.status = Shipped
	return nil
}

func (o *order) Status() Status { return o.status }
func (o *order) Id() string     { return o.ID }
