package pay

import (
	"context"
	"testing"
	"time"

	"github.com/eyazici90/go-ddd/internal/infra/inmem"
	"github.com/eyazici90/go-ddd/pkg/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrder(t *testing.T) {
	orderID := uuid.New().String()

	cmd := OrderCommand{orderID}

	newOrder, err := order.NewOrder(order.OrderID(cmd.OrderID),
		order.NewCustomerID(),
		order.NewProductID(),
		time.Now,
		order.Submitted,
		aggregate.NewVersion())
	require.NoError(t, err)

	handler := NewOrderCommandHandler(orderGetterFunc(func(context.Context, string) (*order.Order, error) {
		return newOrder, nil
	}), inmem.NewOrderRepository())

	err = handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
	assert.Equal(t, order.Paid, newOrder.Status())
}

type orderGetterFunc func(context.Context, string) (*order.Order, error)

func (o orderGetterFunc) Get(ctx context.Context, id string) (*order.Order, error) {
	return o(ctx, id)
}
