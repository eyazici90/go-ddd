package customer

import (
	"github.com/google/uuid"
)

type CustomerId string

func New() CustomerId {
	return CustomerId(uuid.New().String())
}

func (id CustomerId) String() string { return string(id) }
