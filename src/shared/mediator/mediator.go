package mediator

import (
	"context"
	"reflect"
)

type Mediator interface {
	initializer
	sender
	publisher
	pipelineBuilder
}

type reflectBasedMediator struct {
	behaviour    func(context.Context, interface{}) error
	handlers     map[reflect.Type]interface{}
	handlersFunc map[reflect.Type]reflect.Value
}

func NewMediator() Mediator {
	return &reflectBasedMediator{
		handlers:     make(map[reflect.Type]interface{}),
		handlersFunc: make(map[reflect.Type]reflect.Value),
	}
}
