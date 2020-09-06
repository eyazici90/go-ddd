package mediator

import (
	"reflect"
)

type Next func(interface{}) error

type Mediator interface {
	Send(msg interface{}) error
	Publish(msg interface{})
	RegisterHandler(handler interface{}) Mediator
	RegisterBehaviour(func(interface{}, Next) error) Mediator
}

type reflectBasedMediator struct {
	behaviours   func(interface{}, Next) error
	handlers     map[reflect.Type]interface{}
	handlersFunc map[reflect.Type]reflect.Value
}

func New() Mediator {
	return &reflectBasedMediator{
		handlers:     make(map[reflect.Type]interface{}),
		handlersFunc: make(map[reflect.Type]reflect.Value),
	}
}

func (m *reflectBasedMediator) Send(msg interface{}) error {
	if m.behaviours != nil {
		return m.behaviours(msg, m.send)
	}
	return m.send(msg)
}

func (m *reflectBasedMediator) send(msg interface{}) error {
	msgType := reflect.TypeOf(msg)
	handler, _ := m.handlers[msgType]
	handlerFunc, _ := m.handlersFunc[msgType]
	return call(handler, handlerFunc, msg)
}

func (m *reflectBasedMediator) Publish(msg interface{}) {
	// return callHandle(handler, msg)
}

func (m *reflectBasedMediator) RegisterHandler(handler interface{}) Mediator {
	handlerType := reflect.TypeOf(handler)
	method, ok := handlerType.MethodByName(HandleMethodName)
	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType.String())
	}

	cType := reflect.TypeOf(method.Func.Interface()).In(1)

	// fmt.Println(cType)

	m.handlers[cType] = handler
	m.handlersFunc[cType] = method.Func
	return m
}

func (m *reflectBasedMediator) RegisterBehaviour(behaviour func(interface{}, Next) error) Mediator {
	m.behaviours = behaviour
	return m
}
