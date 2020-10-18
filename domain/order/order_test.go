package order

import (
	"orderContext/shared/aggregate"
	"testing"
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/product"

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

	assert.Equal(t, Cancelled, o.Status())
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

	assert.Equal(t, OrderNotPaidError, err)
}

func fakeOrder() *Order {
	o, _ := NewOrder("123", customer.New(), product.New(), func() time.Time { return time.Now() }, Submitted, aggregate.NewVersion())
	return o
}
