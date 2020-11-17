package behaviour

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrier(t *testing.T) {

	retrier := Retry

	retryCount := 0

	var networkError = errors.New("Fake network error")

	next := func(context.Context) error {
		retryCount++
		return networkError
	}

	_ = retrier(nil, nil, next)

	assert.Equal(t, 2, retryCount)
}
