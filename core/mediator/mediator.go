package mediator

import (
	"reflect"
)

type Next func(interface{}) error

type Mediator interface {
	Send(msg interface{}) error
	Publish(msg interface{})
	RegisterHandler(mType reflect.Type, handlerFactory func() interface{}) Mediator
	RegisterBehaviour(func(interface{}, Next) error) Mediator
}

type reflectBasedMediator struct {
	behaviours func(interface{}, Next) error
	handlers   map[reflect.Type]interface{}
}

func New() Mediator {
	return &reflectBasedMediator{handlers: make(map[reflect.Type]interface{})}
}

func (m *reflectBasedMediator) Send(msg interface{}) error {
	if m.behaviours != nil {
		return m.behaviours(msg, m.send)
	}
	return m.send(msg)
}

func (m *reflectBasedMediator) send(msg interface{}) error {
	msgType := reflect.TypeOf(msg)
	handler := m.handlers[msgType]

	handlerType := reflect.TypeOf(handler)

	handleMethod, ok := handlerType.MethodByName("Handle")

	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType.String())
	}

	in := []reflect.Value{reflect.ValueOf(handler), reflect.ValueOf(msg)}

	result := handleMethod.Func.Call(in)
	if result == nil {
		return nil
	}

	if v := result[0].Interface(); v != nil {
		return v.(error)
	}
	return nil
}

func (m *reflectBasedMediator) Publish(msg interface{}) {

}

func (m *reflectBasedMediator) RegisterHandler(mType reflect.Type, handlerFactory func() interface{}) Mediator {
	m.handlers[mType] = handlerFactory()
	return m
}

func (m *reflectBasedMediator) RegisterBehaviour(behaviour func(interface{}, Next) error) Mediator {
	m.behaviours = behaviour
	return m
}
