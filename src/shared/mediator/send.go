package mediator

import (
	"context"
	"reflect"
)

type sender interface {
	Send(context.Context, interface{}) error
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
