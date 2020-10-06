package order

import (
	"errors"
)

var (
	AggregateNotFound = errors.New("Aggregate not found!")

	OrderNotPaidError = errors.New("Order has not paid yet!")

	InvalidValueError = errors.New("Invalid Value!")
)
