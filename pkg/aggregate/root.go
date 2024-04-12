package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type Now func() time.Time

type Version string

func NewVersion() Version { return Version(uuid.New().String()) }

func (v Version) String() string { return string(v) }

type eventRecorder struct {
	events []interface{}
}

func (e *eventRecorder) Record(event interface{}) {
	e.events = append(e.events, event)
}

func (e *eventRecorder) Events() []interface{} { return e.events }

func (e *eventRecorder) Clear() {
	e.events = []interface{}{}
}

type Root struct {
	er eventRecorder
}

func (root *Root) AddEvent(event interface{}) { root.er.Record(event) }

func (root *Root) Clear() { root.er.Clear() }

func (root *Root) Events() []interface{} { return root.er.Events() }
