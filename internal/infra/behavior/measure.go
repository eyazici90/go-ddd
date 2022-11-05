package behavior

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/eyazici90/go-mediator/mediator"
)

func Measure(ctx context.Context, msg mediator.Message, next mediator.Next) error {
	start := time.Now()

	err := next(ctx)

	elapsed := time.Since(start)
	log.Printf("Execution for the command (%s) took %s", reflect.TypeOf(msg), elapsed)

	return err
}
