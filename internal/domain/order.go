package domain

import (
	"time"

	"ordercontext/pkg/aggregate"
)

type OrderID string

type OrderStatus int

const (
	Unknown OrderStatus = iota
	Submitted
	Paid
	Shipped
	Canceled
)

type Order struct {
	aggregate.Root
	id          OrderID
	customerID  CustomerID
	productID   ProductID
	createdTime time.Time
	status      OrderStatus
	version     aggregate.Version
}

func NewOrder(id OrderID, customerID CustomerID,
	productID ProductID, now aggregate.Now,
	status OrderStatus, version aggregate.Version) (*Order, error) {
	o := Order{
		id:          id,
		customerID:  customerID,
		productID:   productID,
		createdTime: now(),
		status:      status,
		version:     version,
	}

	if err := valid(&o); err != nil {
		return nil, err
	}

	o.AddEvent(CreatedEvent{id: string(id)})

	return &o, nil
}

func (o *Order) Pay() {
	o.status = Paid
	o.AddEvent(PaidEvent{id: string(o.id)})
}

func (o *Order) Cancel() {
	o.status = Canceled
	o.AddEvent(CancelledEvent{id: string(o.id)})
}

func (o *Order) Ship() error {
	if o.status != Paid {
		return ErrNotPaid
	}

	o.status = Shipped
	return nil
}

func (o *Order) ID() string { return string(o.id) }

func (o *Order) ProductID() string { return o.productID.String() }

func (o *Order) CustomerID() string { return o.customerID.String() }

func (o *Order) Version() string { return o.version.String() }

func (o *Order) Status() OrderStatus { return o.status }

func valid(o *Order) error {
	if o.id == "" || o.customerID == "" || o.productID == "" {
		return ErrInvalidValue
	}
	return nil
}
