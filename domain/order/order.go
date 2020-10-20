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
	version     aggregate.Version
}

func NewOrder(id OrderId, customerId customer.CustomerId,
	productId product.ProductId, now aggregate.Now,
	status Status, version aggregate.Version) (*Order, error) {
	o := &Order{
		id:          id,
		customerId:  customerId,
		productId:   productId,
		createdTime: now(),
		status:      status,
		version:     version,
	}

	if err := ValidateState(o); err != nil {
		return nil, err
	}

	o.AddEvent(OrderCreatedEvent{id: string(id)})

	return o, nil
}

func ValidateState(o *Order) error {
	if o.id == "" || o.customerId == "" || o.productId == "" {
		return ErrInvalidValue
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
		return ErrOrderNotPaid
	}

	o.status = Shipped
	return nil
}

func (o *Order) ID() string { return string(o.id) }

func (o *Order) ProductId() string { return o.productId.String() }

func (o *Order) CustomerId() string { return o.customerId.String() }

func (o *Order) Version() string { return o.version.String() }

func (o *Order) Status() Status { return o.status }
