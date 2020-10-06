package order

import "github.com/google/uuid"

type OrderId string

func NewOrderId() OrderId {
	return OrderId(uuid.New().String())
}
