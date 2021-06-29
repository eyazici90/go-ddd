package command

import (
	"context"
	"testing"
	"time"

	"ordercontext/internal/domain"
	"ordercontext/internal/infrastructure/store/order"
	"ordercontext/pkg/aggregate"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	handler := NewCreateOrderCommandHandler(order.NewInMemoryRepository().Create)

	orderID := uuid.New().String()

	cmd := CreateOrderCommand{orderID}

	err := handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {
	orderID := uuid.New().String()

	cmd := PayOrderCommand{orderID}

	newOrder, _ := domain.NewOrder(domain.OrderID(cmd.OrderID),
		domain.NewCustomerID(),
		domain.NewProductID(),
		func() time.Time { return time.Now() },
		domain.Submitted,
		aggregate.NewVersion())

	handler := NewPayOrderCommandHandler(func(context.Context, string) (*domain.Order, error) {
		return newOrder, nil
	}, order.NewInMemoryRepository().Update)

	err := handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
	assert.Equal(t, domain.Paid, newOrder.Status())
}
