package behaviour

import (
	"context"

	"github.com/avast/retry-go"
	"github.com/eyazici90/go-mediator"
)

type Retrier struct{}

func NewRetrier() *Retrier { return &Retrier{} }

func (r *Retrier) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {

	err := next(ctx)

	if err != nil {
		err = retry.Do(func() error {
			return next(ctx)
		}, retry.Attempts(1))
	}

	return err
}
