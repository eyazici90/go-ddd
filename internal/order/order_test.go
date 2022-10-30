package order_test

import (
	"testing"
	"time"

	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	o := fakeOrder()

	assert.NotNil(t, o)
}

func TestPayOrder(t *testing.T) {
	o := fakeOrder()

	o.Pay()

	assert.Equal(t, order.Paid, o.Status())
}

func TestCancelOrder(t *testing.T) {
	o := fakeOrder()

	o.Cancel()

	assert.Equal(t, order.Canceled, o.Status())
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

	assert.Equal(t, order.ErrNotPaid, err)
}

func fakeOrder() *order.Order {
	o, _ := order.NewOrder("123", order.NewCustomerID(), order.NewProductID(),
		time.Now, order.Submitted, aggregate.NewVersion())
	return o
}
