package product

import (
	"github.com/google/uuid"
)

type ProductId string

func New() ProductId {
	return ProductId(uuid.New().String())
}

func (id ProductId) String() string { return string(id) }
