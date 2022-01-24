package event

type Publisher interface {
	Publish(event interface{})
	PublishAll(events ...interface{})
}
