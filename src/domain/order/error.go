package order

import (
	"errors"
)

var (
	OrderNotPaidError = errors.New("Order has not paid yet!")

	InvalidValueError = errors.New("Invalid Value!")
)
