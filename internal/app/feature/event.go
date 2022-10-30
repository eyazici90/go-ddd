package feature

type EventPublisher interface {
	Publish(event interface{})
	PublishAll(events ...interface{})
}
