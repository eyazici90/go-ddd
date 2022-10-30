package behavior

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrier(t *testing.T) {
	retrier := Retry

	retryCount := 0

	networkError := errors.New("fake network error")

	next := func(context.Context) error {
		retryCount++
		return networkError
	}

	_ = retrier(nil, nil, next)

	assert.Equal(t, 2, retryCount)
}
