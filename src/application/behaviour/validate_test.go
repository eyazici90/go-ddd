package behaviour

import (
	"context"
	"orderContext/application/command"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	invalidId := "123"

	cmd := command.CreateOrderCommand{Id: invalidId}

	next := func(context.Context) error {
		return nil
	}
	validator := NewValidator()

	err := validator.Process(nil, cmd, next)

	assert.NotNil(t, err)
}
