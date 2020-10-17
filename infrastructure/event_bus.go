package infrastructure

import (
	"fmt"
	"reflect"
)

//

type EventPublisher interface {
	Publish(event interface{})
	PublishAll(events ...interface{})
}

type NoBus struct{}

func NewNoBus() *NoBus {
	return &NoBus{}
}

func (r *NoBus) Publish(event interface{}) {
	//
	fmt.Println("event that is published :" + reflect.TypeOf(event).Name())
}

func (r *NoBus) PublishAll(events ...interface{}) {
	for _, event := range events {
		r.Publish(event)
	}
}
