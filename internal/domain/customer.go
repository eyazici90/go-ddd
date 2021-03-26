package domain

import (
	"github.com/google/uuid"
)

type CustomerID string

func NewCustomerID() CustomerID {
	return CustomerID(uuid.New().String())
}

func (id CustomerID) String() string { return string(id) }
