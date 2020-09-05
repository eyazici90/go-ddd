package behaviour

import (
	"fmt"
	"orderContext/core/mediator"
)

type LogBehaviour struct{}

func NewLogBehaviour() LogBehaviour { return LogBehaviour{} }

func (l LogBehaviour) Process(cmd interface{}, next mediator.Next) error {

	fmt.Println("Pre Process!")
	result := next(cmd)
	fmt.Println("Post Process")

	return result
}
