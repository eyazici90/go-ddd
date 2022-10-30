package create

import (
	"context"
	"testing"

	"github.com/eyazici90/go-ddd/internal/infra/inmem"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	handler := NewCreateOrderHandler(inmem.NewOrderRepository())

	orderID := uuid.New().String()

	cmd := CreateOrder{orderID}

	err := handler.Handle(context.TODO(), cmd)

	assert.Nil(t, err)
}
