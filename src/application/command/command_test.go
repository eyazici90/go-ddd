package command

import (
	"orderContext/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCreateOrder(t *testing.T) {
	handler := NewCreateOrderCommandHandler(infrastructure.NewOrderRepository())

	orderId := uuid.New().String()

	cmd := CreateOrderCommand{orderId}

	err := handler.Handle(nil, cmd)

	assert.Nil(t, err)
}
