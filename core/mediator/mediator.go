package mediator

import (
	"context"
	"reflect"
)

type Mediator interface {
	Send(context.Context, interface{}) error
	Publish(msg interface{})
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

func (m *reflectBasedMediator) Send(ctx context.Context, msg interface{}) error {
	if m.behaviour != nil {
		return m.behaviour(ctx, msg)
	}
	return m.send(ctx, msg)
}

func (m *reflectBasedMediator) send(ctx context.Context, msg interface{}) error {
	msgType := reflect.TypeOf(msg)
	handler, _ := m.handlers[msgType]
	handlerFunc, _ := m.handlersFunc[msgType]
	return call(handler, ctx, handlerFunc, msg)
}

func (m *reflectBasedMediator) Publish(msg interface{}) {
}

func (m *reflectBasedMediator) RegisterHandler(handler interface{}) Mediator {
	handlerType := reflect.TypeOf(handler)

	method, ok := handlerType.MethodByName(HandleMethodName)
	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType.String())
	}

	cType := reflect.TypeOf(method.Func.Interface()).In(2)

	m.handlers[cType] = handler
	m.handlersFunc[cType] = method.Func
	return m
}

func (m *reflectBasedMediator) UseBehaviour(pipelineBehaviour PipelineBehaviour) Mediator {
	return m.Use(pipelineBehaviour.Process)
}

func (m *reflectBasedMediator) Use(call func(context.Context, interface{}, Next) error) Mediator {
	if m.behaviour == nil {
		m.behaviour = m.send
	}
	seed := m.behaviour

	m.behaviour = func(ctx context.Context, msg interface{}) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}

	return m
}
