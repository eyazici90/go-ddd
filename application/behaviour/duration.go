package behaviour

import (
	"context"
	"orderContext/core/mediator"
	"time"
)

type Cancellator struct {
}

func NewCancellator() *Cancellator { return &Cancellator{} }

func (l *Cancellator) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {
	c, cancel := context.WithTimeout(ctx, time.Duration(60*time.Second))
	defer cancel()

	result := next(c)

	return result
}
