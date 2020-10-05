package behaviour

import (
	"context"
	"log"

	"github.com/eyazici90/go-mediator"
)

type Logger struct{}

func NewLogger() *Logger { return &Logger{} }

func (l *Logger) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {

	log.Println("Pre Process of the command!")

	if err := next(ctx); err != nil {
		return err
	}

	log.Println("Post Process of the command")

	return nil
}
