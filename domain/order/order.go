package order

import (
	"orderContext/domain/customer"
	"orderContext/domain/product"
)

type Status int

const (
	Submitted Status = 1
	Paid      Status = 2
	Shipped   Status = 3
	Cancelled Status = 4
)

type OrderId string

type Order struct {
	id         OrderId
	customerId customer.CustomerId
	productId  product.ProductId
	status     Status
}

func NewOrder(id OrderId, customerId customer.CustomerId, productId product.ProductId) *Order {
	return &Order{
		id:         id,
		customerId: customerId,
		productId:  productId,
	}
}

func (o *Order) Pay() { o.status = Paid }

func (o *Order) Cancel() { o.status = Cancelled }

func (o *Order) Ship() error {

	if o.status != Paid {
		return OrderNotPaidError
	}

	o.status = Shipped

	return nil
}
