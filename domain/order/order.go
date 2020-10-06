package order

import (
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/product"
	"orderContext/shared/aggregate"
)

type Order interface {
	aggregate.EventRecorder
	Pay()
	Cancel()
	Ship() error
	Status() Status
	Id() string
}

type order struct {
	aggregate.AggregateRoot
	id          OrderId
	customerId  customer.CustomerId
	productId   product.ProductId
	createdTime time.Time
	status      Status
}

func NewOrder(id OrderId, customerId customer.CustomerId, productId product.ProductId, now aggregate.Now) (Order, error) {
	o := &order{
		customerId:  customerId,
		productId:   productId,
		createdTime: now(),
	}
	o.id = id

	err := ValidateState(o)

	if err != nil {
		return nil, err
	}

	o.AddEvent(OrderCreatedEvent{id: string(id)})

	return o, nil
}

func ValidateState(o *order) error {
	if o.id == "" || o.customerId == "" || o.productId == "" {
		return InvalidValueError
	}
	return nil
}

func (o *order) Pay() {
	o.status = Paid
	o.AddEvent(OrderPaidEvent{id: string(o.id)})
}

func (o *order) Cancel() {
	o.status = Cancelled
	o.AddEvent(OrderCancelledEvent{id: string(o.id)})
}

func (o *order) Ship() error {

	if o.status != Paid {
		return OrderNotPaidError
	}

	o.status = Shipped
	return nil
}

func (o *order) Status() Status { return o.status }
func (o *order) Id() string     { return string(o.id) }
