package mediator

import "context"

type PipelineBehaviour interface {
	Process(context.Context, interface{}, Next) error
}
