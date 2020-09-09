package order

import (
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

	assert.Equal(t, Paid, o.status)
}

func TestCancelOrder(t *testing.T) {
	o := fakeOrder()

	o.Cancel()

	assert.Equal(t, Cancelled, o.status)
}

func TestShipOrder(t *testing.T) {
	o := fakeOrder()
	o.Pay()

	err := o.Ship()

	assert.Nil(t, err)
}

func fakeOrder() *Order {
	return NewOrder("123", customer.New(), product.New(), func() time.Time { return time.Now() })
}
