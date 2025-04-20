package order

import (
	"github.com/google/uuid"
)

type ProductID string

func NewProductID() ProductID {
	return ProductID(uuid.New().String())
}

func (id ProductID) String() string { return string(id) }
