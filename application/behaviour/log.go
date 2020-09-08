package behaviour

import (
	"context"
	"fmt"
	"orderContext/core/mediator"
)

type Logger struct {
	next mediator.Next
}

func NewLogger() *Logger { return &Logger{} }

func (l *Logger) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {

	fmt.Println("Pre Process!")

	result := next()

	fmt.Println("Post Process")

	return result
}