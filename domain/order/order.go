package order

import (
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/product"
	"orderContext/shared/aggregate"
)

type Order struct {
	aggregate.AggregateRoot
	id          OrderId
	customerId  customer.CustomerId
	productId   product.ProductId
	createdTime time.Time
	status      Status
}

func NewOrder(id OrderId, customerId customer.CustomerId, productId product.ProductId, now aggregate.Now) (*Order, error) {
	o := &Order{
		id:          id,
		customerId:  customerId,
		productId:   productId,
		createdTime: now(),
		status:      Submitted,
	}

	if err := ValidateState(o); err != nil {
		return nil, err
	}

	o.AddEvent(OrderCreatedEvent{id: string(id)})

	return o, nil
}

func ValidateState(o *Order) error {
	if o.id == "" || o.customerId == "" || o.productId == "" {
		return InvalidValueError
	}
	return nil
}

func (o *Order) Pay() {
	o.status = Paid
	o.AddEvent(OrderPaidEvent{id: string(o.id)})
}

func (o *Order) Cancel() {
	o.status = Cancelled
	o.AddEvent(OrderCancelledEvent{id: string(o.id)})
}

func (o *Order) Ship() error {

	if o.status != Paid {
		return OrderNotPaidError
	}

	o.status = Shipped
	return nil
}

func (o *Order) Status() Status { return o.status }
func (o *Order) Id() string     { return string(o.id) }
