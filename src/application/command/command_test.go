package command

import (
	"context"
	"orderContext/domain/customer"
	"orderContext/domain/order"
	"orderContext/domain/product"
	"orderContext/infrastructure"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCreateOrder(t *testing.T) {
	handler := NewCreateOrderCommandHandler(infrastructure.NewOrderRepository().Create)

	orderId := uuid.New().String()

	cmd := CreateOrderCommand{orderId}

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {

	orderId := uuid.New().String()

	cmd := PayOrderCommand{orderId}

	newOrder, _ := order.NewOrder(order.OrderId(cmd.OrderId), customer.New(), product.New(), func() time.Time { return time.Now() })

	handler := NewPayOrderCommandHandler(func(context.Context, string) order.Order {
		return newOrder
	}, infrastructure.NewOrderRepository().Update)

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
	assert.Equal(t, order.Paid, newOrder.Status())

}
