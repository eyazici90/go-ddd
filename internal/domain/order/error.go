package order

import (
	"errors"
)

var (
	ErrOrderNotPaid = errors.New("order has not paid yet")
	ErrInvalidValue = errors.New("invalid value")
)
