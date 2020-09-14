package infrastructure

//

type EventPublisher interface {
	Publish(event interface{})
	PublishAll(events ...interface{})
}

type RabbitMQBus struct{}

func (r *RabbitMQBus) Publish(event interface{}) {
	//
}

func (r *RabbitMQBus) PublishAll(events ...interface{}) {
	for _, event := range events {
		r.Publish(event)
	}
}
