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

type RabbitMQBus struct{}

func NewRabbitMQBus() *RabbitMQBus {
	return &RabbitMQBus{}
}

func (r *RabbitMQBus) Publish(event interface{}) {
	//
	fmt.Println("event that is published :" + reflect.ValueOf(event).String())
}

func (r *RabbitMQBus) PublishAll(events ...interface{}) {
	for _, event := range events {
		r.Publish(event)
	}
}
