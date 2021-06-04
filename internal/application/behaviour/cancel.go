package behaviour

import (
	"context"

	"time"

	"github.com/eyazici90/go-mediator/mediator"
)

type Cancellator struct {
	timeout time.Duration
}

func NewCancellator(timeout time.Duration) *Cancellator { return &Cancellator{timeout} }

func (c *Cancellator) Process(ctx context.Context, _ mediator.Message, next mediator.Next) error {

	timeoutContext, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return next(timeoutContext)
}
