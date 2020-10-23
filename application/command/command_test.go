package command

import (
	"context"
	"ordercontext/domain/customer"
	"ordercontext/domain/order"
	"ordercontext/domain/product"
	"ordercontext/infrastructure"
	"ordercontext/shared/aggregate"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCreateOrder(t *testing.T) {
	handler := NewCreateOrderCommandHandler(infrastructure.InMemoryRepository.Create)

	orderId := uuid.New().String()

	cmd := CreateOrderCommand{orderId}

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {

	orderId := uuid.New().String()

	cmd := PayOrderCommand{orderId}

	newOrder, _ := order.NewOrder(order.OrderId(cmd.OrderId), customer.New(), product.New(), func() time.Time { return time.Now() }, order.Submitted, aggregate.NewVersion())

	handler := NewPayOrderCommandHandler(func(context.Context, string) (*order.Order, error) {
		return newOrder, nil
	}, infrastructure.InMemoryRepository.Update)

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
	assert.Equal(t, order.Paid, newOrder.Status())

}
