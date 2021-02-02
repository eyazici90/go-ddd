package aggregate

type Root struct {
	eventRecorder eventRecorder
}

func (root *Root) AddEvent(event interface{}) { root.eventRecorder.Record(event) }

func (root *Root) Clear() { root.eventRecorder.Clear() }

func (root *Root) Events() []interface{} { return root.eventRecorder.Events() }
