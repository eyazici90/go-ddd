package aggregate

type EventRecorder interface {
	AddEvent(event interface{})
	Events() []interface{}
	Clear()
}

type AggregateRoot struct {
	eventRecorder eventRecorder
}

func (root *AggregateRoot) AddEvent(event interface{}) { root.eventRecorder.Record(event) }

func (root *AggregateRoot) Clear() { root.eventRecorder.Clear() }

func (root *AggregateRoot) Events() []interface{} { return root.eventRecorder.Events() }
