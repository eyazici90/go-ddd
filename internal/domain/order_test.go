package domain

import (
	"testing"
	"time"

	"ordercontext/pkg/aggregate"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	o := fakeOrder()

	assert.NotNil(t, o)
}

func TestPayOrder(t *testing.T) {
	o := fakeOrder()

	o.Pay()

	assert.Equal(t, Paid, o.Status())
}

func TestCancelOrder(t *testing.T) {
	o := fakeOrder()

	o.Cancel()

	assert.Equal(t, Canceled, o.Status())
}

func TestShipOrder(t *testing.T) {
	o := fakeOrder()

	o.Pay()

	err := o.Ship()

	assert.Nil(t, err)
}

func TestShipOrderWithoutPaidExpectErr(t *testing.T) {
	o := fakeOrder()

	err := o.Ship()

	assert.Equal(t, ErrOrderNotPaid, err)
}

func fakeOrder() *Order {
	o, _ := NewOrder("123", NewCustomerID(), NewProductID(),
		func() time.Time { return time.Now() }, Submitted, aggregate.NewVersion())
	return o
}
