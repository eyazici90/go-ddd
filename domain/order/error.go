package order

import (
	"errors"
)

var (
	AggregateNotFound = errors.New("aggregate not found!")

	OrderNotPaidError = errors.New("order has not paid yet!")

	InvalidValueError = errors.New("invalid value!")
)
