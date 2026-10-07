package scheduler

// Drain returns at most max events, state events first, then the newest
// progress per job.
func (o *Outbox) Drain(max int) []Event {
	if max <= 0 {
		return nil
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	out := o.takeStates(max)
	out = append(out, o.takeProgress(max-len(out))...)

	return out
}

// takeStates pops up to max state events.
func (o *Outbox) takeStates(max int) []Event {
	if max > len(o.states) {
		max = len(o.states)
	}

	out := make([]Event, max)
	copy(out, o.states[:max])
	o.states = o.states[max:]

	return out
}

// takeProgress pops up to max coalesced progress events in first-seen order.
func (o *Outbox) takeProgress(max int) []Event {
	var out []Event

	for len(o.order) > 0 && len(out) < max {
		id := o.order[0]
		o.order = o.order[1:]
		out = append(out, o.latest[id])
		delete(o.latest, id)
	}

	return out
}

// Updated returns a signal channel that fires when events arrive.
func (o *Outbox) Updated() <-chan struct{} { return o.updated }
