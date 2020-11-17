package behaviour

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/eyazici90/go-mediator"
)

func Measure(ctx context.Context, cmd mediator.Message, next mediator.Next) error {
	start := time.Now()

	err := next(ctx)

	elapsed := time.Since(start)

	log.Printf("Execution for the command (%s) took %s", reflect.TypeOf(cmd), elapsed)

	return err
}
