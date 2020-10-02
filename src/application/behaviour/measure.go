package behaviour

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/eyazici90/go-mediator"
)

type Measurer struct{}

func NewMeasurer() *Measurer { return &Measurer{} }

func (c *Measurer) Process(ctx context.Context, cmd interface{}, next mediator.Next) error {
	start := time.Now()

	result := next(ctx)

	elapsed := time.Since(start)

	log.Printf("Execution for the command (%s) took %s", reflect.TypeOf(cmd), elapsed)

	return result
}
