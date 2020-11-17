package order

import (
	"errors"
)

var (
	ErrAggregateNotFound = errors.New("aggregate not found")

	ErrOrderNotPaid = errors.New("order has not paid yet")

	ErrInvalidValue = errors.New("invalid value")
)
