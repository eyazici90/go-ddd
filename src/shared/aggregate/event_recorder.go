package aggregate

type EventRecorder struct {
	events []interface{}
}

func (e *EventRecorder) Record(event interface{}) {
	e.events = append(e.events, event)
}

func (e *EventRecorder) Events() []interface{} { return e.events }

func (e *EventRecorder) Clear() {
	e.events = []interface{}{}
}
