package shared

type AggregateRoot struct {
	events []interface{}
}

func (a *AggregateRoot) AddEvent(event interface{}) {
	a.events = append(a.events, event)
}

func (a *AggregateRoot) Events() []interface{} { return a.events }

func (a *AggregateRoot) ClearEvents() {
	a.events = []interface{}{}
}
