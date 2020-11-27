package order

import (
	"time"

	"ordercontext/internal/domain/customer"
	"ordercontext/internal/domain/product"
	"ordercontext/pkg/aggregate"

	"github.com/google/uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

type Order struct {
	aggregate.Root
	id          ID
	customerID  customer.ID
	productID   product.ID
	createdTime time.Time
	status      Status
	version     aggregate.Version
}

func NewOrder(id ID, customerId customer.ID,
	productId product.ID, now aggregate.Now,
	status Status, version aggregate.Version) (*Order, error) {
	o := &Order{
		id:          id,
		customerID:  customerId,
		productID:   productId,
		createdTime: now(),
		status:      status,
		version:     version,
	}

	if err := ValidateState(o); err != nil {
		return nil, err
	}

	o.AddEvent(CreatedEvent{id: string(id)})

	return o, nil
}

func ValidateState(o *Order) error {
	if o.id == "" || o.customerID == "" || o.productID == "" {
		return ErrInvalidValue
	}
	return nil
}

func (o *Order) Pay() {
	o.status = Paid
	o.AddEvent(PaidEvent{id: string(o.id)})
}

func (o *Order) Cancel() {
	o.status = Cancelled
	o.AddEvent(CancelledEvent{id: string(o.id)})
}

func (o *Order) Ship() error {

	if o.status != Paid {
		return ErrOrderNotPaid
	}

	o.status = Shipped
	return nil
}

func (o *Order) ID() string { return string(o.id) }

func (o *Order) ProductID() string { return o.productID.String() }

func (o *Order) CustomerID() string { return o.customerID.String() }

func (o *Order) Version() string { return o.version.String() }

func (o *Order) Status() Status { return o.status }
