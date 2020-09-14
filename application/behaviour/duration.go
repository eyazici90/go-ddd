package behaviour

import (
	"context"
	"orderContext/core/mediator"
	"time"

	"github.com/spf13/viper"
)

type Cancellator struct {
}

func NewCancellator() *Cancellator { return &Cancellator{} }

func (l *Cancellator) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {
	timeout := viper.GetInt("context.timeout")
	c, cancel := context.WithTimeout(ctx, time.Duration(time.Duration(timeout)*time.Second))
	defer cancel()

	result := next(c)

	return result
}
