package wire

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/scheduler"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// Poll returns up to budget results without waiting: engine events first,
// then worker answers.
func (b *Backend) Poll(budget int) []ui.Update {
	if budget <= 0 {
		return nil
	}

	out := make([]ui.Update, 0, budget)
	for _, ev := range b.eng.Drain(budget) {
		out = append(out, b.event(ev))
	}

	for len(out) < budget {
		select {
		case u := <-b.results:
			out = append(out, u)
		default:
			return out
		}
	}

	return out
}

// event maps one engine event. Progress events feed the speed tracker; a
// state event carries the full job, including a terminal state.
func (b *Backend) event(ev scheduler.Event) ui.Update {
	if ev.Kind == scheduler.EventProgress {
		b.track(ev)
	}

	row := b.mapJob(ev.Job)

	return ui.Update{Kind: ui.UpdateProgress, Row: &row}
}

// track folds one progress event into the per-job speed. The first event
// only records the baseline; later events smooth the rate.
func (b *Backend) track(ev scheduler.Event) {
	at := ev.At
	if at.IsZero() {
		at = time.Now()
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	prev := b.speeds[ev.Job.ID]
	if !prev.at.IsZero() && at.After(prev.at) {
		delta := ev.Job.Done - prev.done
		if delta >= 0 {
			inst := float64(delta) / at.Sub(prev.at).Seconds()
			if prev.bytes <= 0 {
				prev.bytes = inst
			} else {
				prev.bytes = 0.7*prev.bytes + 0.3*inst
			}
		}
	}

	prev.done, prev.at = ev.Job.Done, at
	b.speeds[ev.Job.ID] = prev
}
