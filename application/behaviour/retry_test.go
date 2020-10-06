package behaviour

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrier(t *testing.T) {

	retrier := NewRetrier()

	retryCount := 0

	var networkError = errors.New("Fake network error")

	next := func(context.Context) error {
		retryCount++
		return networkError
	}

	_ = retrier.Process(nil, nil, next)

	assert.Equal(t, 2, retryCount)
}
