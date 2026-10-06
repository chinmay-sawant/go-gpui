package scheduler

import (
	"sync"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// Kind labels an event.
type Kind int

// The two event kinds. Progress may coalesce; State never drops.
const (
	EventProgress Kind = iota
	EventState
)

// Event is one observation for the UI. A State event carries the whole job,
// including Error on failure; a Progress event carries the newest bytes.
type Event struct {
	Kind Kind
	Job  domain.Job
	At   time.Time
}

// Outbox coalesces progress per job and keeps every state event in order.
// Drain never blocks the producer.
type Outbox struct {
	mu      sync.Mutex
	latest  map[string]Event
	order   []string
	states  []Event
	updated chan struct{}
}

// NewOutbox builds an empty outbox.
func NewOutbox() *Outbox {
	return &Outbox{
		latest:  map[string]Event{},
		updated: make(chan struct{}, 1),
	}
}

// publish stores one event. A progress event replaces the previous progress
// for the same job; a state event is appended behind existing state events.
func (o *Outbox) publish(ev Event) {
	o.mu.Lock()
	if ev.Kind == EventProgress {
		if _, seen := o.latest[ev.Job.ID]; !seen {
			o.order = append(o.order, ev.Job.ID)
		}

		o.latest[ev.Job.ID] = ev
	} else {
		o.states = append(o.states, ev)
	}
	o.mu.Unlock()

	select {
	case o.updated <- struct{}{}:
	default:
	}
}
