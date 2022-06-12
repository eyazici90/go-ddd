package domain_test

import (
	"testing"
	"time"

	"github.com/eyazici90/go-ddd/internal/domain"
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

	assert.Equal(t, domain.Paid, o.Status())
}

func TestCancelOrder(t *testing.T) {
	o := fakeOrder()

	o.Cancel()

	assert.Equal(t, domain.Canceled, o.Status())
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

	assert.Equal(t, domain.ErrNotPaid, err)
}

func fakeOrder() *domain.Order {
	o, _ := domain.NewOrder("123", domain.NewCustomerID(), domain.NewProductID(),
		time.Now, domain.Submitted, aggregate.NewVersion())
	return o
}
