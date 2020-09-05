package aggregate

type AggregateRoot struct {
	ID            string
	eventRecorder EventRecorder
}

func (root *AggregateRoot) AddEvent(event interface{}) { root.eventRecorder.Record(event) }

func (root *AggregateRoot) Clear() { root.eventRecorder.Clear() }
