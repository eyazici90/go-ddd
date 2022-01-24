package behavior

import (
	"context"

	"github.com/avast/retry-go"

	"github.com/eyazici90/go-mediator/pkg/mediator"
)

func Retry(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	err := retry.Do(func() error {
		return next(ctx)
	}, retry.Attempts(2))

	return err
}
