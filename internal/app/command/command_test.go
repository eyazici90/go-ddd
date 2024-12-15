package command_test

import (
	"context"
	"github.com/eyazici90/go-ddd/pkg/otel"
	"os"
	"testing"
	"time"

	"github.com/eyazici90/go-ddd/internal/app/command"
	"github.com/eyazici90/go-ddd/internal/domain"
	"github.com/eyazici90/go-ddd/internal/infra/inmem"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	handler := command.NewCreateOrderHandler(inmem.NewOrderRepository())

	orderID := uuid.New().String()

	cmd := command.CreateOrder{orderID}

	err := handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
}

func TestPayOrder(t *testing.T) {
	orderID := uuid.New().String()

	cmd := command.PayOrder{orderID}

	newOrder, err := domain.NewOrder(domain.OrderID(cmd.OrderID),
		domain.NewCustomerID(),
		domain.NewProductID(),
		time.Now,
		domain.Submitted,
		aggregate.NewVersion())
	require.NoError(t, err)

	handler := command.NewPayOrderHandler(orderGetterFunc(func(context.Context, string) (*domain.Order, error) {
		return newOrder, nil
	}), inmem.NewOrderRepository())

	err = handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
	assert.Equal(t, domain.Paid, newOrder.Status())
}

type orderGetterFunc func(context.Context, string) (*domain.Order, error)

func (o orderGetterFunc) Get(ctx context.Context, id string) (*domain.Order, error) {
	return o(ctx, id)
}

func TestMain(m *testing.M) {
	_, _ = otel.New(context.Background(), &otel.Config{})
	code := m.Run()
	os.Exit(code)
}
