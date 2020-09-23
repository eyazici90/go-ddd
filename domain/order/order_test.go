package order

import (
	"testing"
	"time"

	"orderContext/domain/customer"
	"orderContext/domain/product"

	"github.com/stretchr/testify/assert"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Given order factory", t, func() {
		factory := func() Order { return fakeOrder() }

		Convey("When called", func() {
			order := factory()

			Convey("Then it should not be null", func() {
				So(order, ShouldNotBeNil)
			})
		})
	})

	Convey("Given order", t, func() {
		order := fakeOrder()

		Convey("When order is paid", func() {
			order.Pay()

			Convey("Then status should be Paid", func() {
				So(order.Status(), ShouldEqual, Paid)
			})
		})

		Convey("When order is cancelled", func() {
			order.Cancel()

			Convey("Then status should be Cancelled", func() {
				So(order.Status(), ShouldEqual, Cancelled)
			})
		})

	})

}

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

func fakeOrder() Order {
	o, _ := NewOrder("123", customer.New(), product.New(), func() time.Time { return time.Now() })
	return o
}
