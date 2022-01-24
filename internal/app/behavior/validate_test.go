package behavior

import (
	"context"
	"testing"

	"github.com/eyazici90/go-ddd/internal/app/command"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	invalidID := "123"

	cmd := command.CreateOrder{ID: invalidID}

	next := func(context.Context) error {
		return nil
	}
	validator := Validate

	err := validator(nil, cmd, next)

	assert.NotNil(t, err)
}
