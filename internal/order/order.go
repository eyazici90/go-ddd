package order

import (
	"time"

	"github.com/eyazici90/go-ddd/pkg/aggregate"
)

type ID string

type Status int

const (
	Unknown Status = iota
	Submitted
	Paid
	Shipped
	Canceled
)

type Order struct {
	aggregate.Root

	id          ID
	customerID  CustomerID
	productID   ProductID
	createdTime time.Time
	status      Status
	version     aggregate.Version
}

func New(id ID, customerID CustomerID,
	productID ProductID, now aggregate.Now,
	status Status, version aggregate.Version,
) (*Order, error) {
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

func (o *Order) Status() Status { return o.status }

func valid(o *Order) error {
	if o.id == "" || o.customerID == "" || o.productID == "" {
		return ErrInvalidValue
	}
	return nil
}
