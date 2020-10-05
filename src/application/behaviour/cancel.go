package behaviour

import (
	"context"

	"time"

	"github.com/eyazici90/go-mediator"
)

type Cancellator struct {
	timeout int
}

func NewCancellator(timeout int) *Cancellator { return &Cancellator{timeout} }

func (c *Cancellator) Process(ctx context.Context, _ interface{}, next mediator.Next) error {

	timeoutContext, cancel := context.WithTimeout(ctx, time.Duration(time.Duration(c.timeout)*time.Second))
	defer cancel()

	return next(timeoutContext)
}
