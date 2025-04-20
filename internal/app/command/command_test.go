package command_test

import (
	"context"
	"github.com/eyazici90/go-ddd/pkg/otel"
	"os"
	"testing"
	"time"

	"github.com/eyazici90/go-ddd/internal/app/command"
	"github.com/eyazici90/go-ddd/internal/infra/mem"
	"github.com/eyazici90/go-ddd/internal/order"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	handler := command.NewCreateOrderHandler(mem.NewOrderRepository())

	orderID := uuid.New().String()

	cmd := command.CreateOrder{orderID}

	err := handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {
	orderID := uuid.New().String()

	cmd := command.PayOrder{orderID}

	newOrder, err := order.New(order.ID(cmd.OrderID),
		order.NewCustomerID(),
		order.NewProductID(),
		time.Now,
		order.Submitted,
		aggregate.NewVersion())
	require.NoError(t, err)

	handler := command.NewPayOrderHandler(orderGetterFunc(func(context.Context, string) (*order.Order, error) {
		return newOrder, nil
	}), mem.NewOrderRepository())

	err = handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
	assert.Equal(t, order.Paid, newOrder.Status())
}

type orderGetterFunc func(context.Context, string) (*order.Order, error)

func (o orderGetterFunc) Get(ctx context.Context, id string) (*order.Order, error) {
	return o(ctx, id)
}

func TestMain(m *testing.M) {
	_, _ = otel.New(context.Background(), &otel.Config{})
	code := m.Run()
	os.Exit(code)
}
