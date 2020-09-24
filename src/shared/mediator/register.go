package mediator

import "reflect"

type initializer interface {
	RegisterHandler(handler interface{}) Mediator
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
