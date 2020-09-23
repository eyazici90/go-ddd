package mediator

import (
	"context"
	"reflect"
)

type Mediator interface {
	sender
	publisher
	RegisterHandler(interface{}) Mediator
	UseBehaviour(PipelineBehaviour) Mediator
	Use(call func(context.Context, interface{}, Next) error) Mediator
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

func (m *reflectBasedMediator) RegisterHandler(handler interface{}) Mediator {
	handlerType := reflect.TypeOf(handler)

	method, ok := handlerType.MethodByName(handleMethodName)
	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType.String())
	}

	cType := reflect.TypeOf(method.Func.Interface()).In(2)

	m.handlers[cType] = handler
	m.handlersFunc[cType] = method.Func
	return m
}
