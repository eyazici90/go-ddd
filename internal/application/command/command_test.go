package command_test

import (
	"context"
	"testing"
	"time"

	"ordercontext/internal/application/command"
	"ordercontext/internal/domain/order"
	"ordercontext/internal/infra/store"
	"ordercontext/pkg/aggregate"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	handler := command.NewCreateOrderCommandHandler(store.NewOrderInMemoryRepository().Create)

	orderID := uuid.New().String()

	cmd := command.CreateOrderCommand{orderID}

	err := handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {
	orderID := uuid.New().String()

	cmd := command.PayOrderCommand{orderID}

	newOrder, err := order.NewOrder(order.OrderID(cmd.OrderID),
		order.NewCustomerID(),
		order.NewProductID(),
		func() time.Time { return time.Now() },
		order.Submitted,
		aggregate.NewVersion())
	require.NoError(t, err)

	handler := command.NewPayOrderCommandHandler(func(context.Context, string) (*order.Order, error) {
		return newOrder, nil
	}, store.NewOrderInMemoryRepository().Update)

	err = handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
	assert.Equal(t, order.Paid, newOrder.Status())
}
