package order_test

import (
	"testing"
	"time"

	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/stretchr/testify/assert"
)

func TestOrderNew(t *testing.T) {
	o := fakeOrder()

	assert.NotNil(t, o)
}

func TestOrderPay(t *testing.T) {
	o := fakeOrder()

	o.Pay()

	assert.Equal(t, order.Paid, o.Status())
}

func TestOrderCancel(t *testing.T) {
	o := fakeOrder()

	o.Cancel()

	assert.Equal(t, order.Canceled, o.Status())
}

func TestOrderShip(t *testing.T) {
	o := fakeOrder()

	o.Pay()

	err := o.Ship()

	assert.Nil(t, err)
}

func TestOrderShipWithoutPaidExpectErr(t *testing.T) {
	o := fakeOrder()

	err := o.Ship()

	assert.Equal(t, order.ErrNotPaid, err)
}

func fakeOrder() *order.Order {
	o, _ := order.New("123", order.NewCustomerID(), order.NewProductID(),
		time.Now, order.Submitted, aggregate.NewVersion())
	return o
}
