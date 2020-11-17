package order

import "github.com/google/uuid"

type OrderID string

func NewOrderID() OrderID {
	return OrderID(uuid.New().String())
}
