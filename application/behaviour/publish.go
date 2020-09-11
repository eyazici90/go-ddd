package behaviour

import (
	"context"
	"orderContext/core/mediator"
)

type Publisher struct {
}

func NewPublisher() *Publisher { return &Publisher{} }

func (l *Publisher) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {

	result := next(ctx)

	return result
}
