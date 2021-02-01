package command

import (
	"context"
	"testing"
	"time"

	"ordercontext/internal/domain/customer"
	"ordercontext/internal/domain/order"
	"ordercontext/internal/domain/product"
	store "ordercontext/internal/infrastructure/store/order"
	"ordercontext/pkg/aggregate"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	handler := NewCreateOrderCommandHandler(store.NewInMemoryRepository().Create)

	orderId := uuid.New().String()

	cmd := CreateOrderCommand{orderId}

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {

	orderId := uuid.New().String()

	cmd := PayOrderCommand{orderId}

	newOrder, _ := order.NewOrder(order.ID(cmd.OrderID), customer.New(), product.New(), func() time.Time { return time.Now() }, order.Submitted, aggregate.NewVersion())

	handler := NewPayOrderCommandHandler(func(context.Context, string) (*order.Order, error) {
		return newOrder, nil
	}, store.NewInMemoryRepository().Update)

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
	assert.Equal(t, order.Paid, newOrder.Status())
}
