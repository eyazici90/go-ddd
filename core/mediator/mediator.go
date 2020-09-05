package mediator

import (
	"reflect"
)

type Mediator interface {
	Send(msg interface{})
	Publish(msg interface{})
	RegisterHandler(mType reflect.Type, handlerFactory func() interface{}) Mediator
}

type reflectBasedMediator struct {
	behaviours []func(interface{}, interface{})
	handlers   map[reflect.Type]interface{}
}

func New() Mediator {
	return reflectBasedMediator{handlers: make(map[reflect.Type]interface{})}
}

func (m reflectBasedMediator) Send(msg interface{}) {
	msgType := reflect.TypeOf(msg)
	handler := m.handlers[msgType]

	handlerType := reflect.TypeOf(handler)

	handleMethod, ok := handlerType.MethodByName("Handle")

	if !ok {
		panic("handle method does not exists")
	}

	in := []reflect.Value{reflect.ValueOf(handler), reflect.ValueOf(msg)}

	handleMethod.Func.Call(in)
}

func (m reflectBasedMediator) Publish(msg interface{}) {

}

func (m reflectBasedMediator) RegisterHandler(mType reflect.Type, handlerFactory func() interface{}) Mediator {
	m.handlers[mType] = handlerFactory()
	return m
}
