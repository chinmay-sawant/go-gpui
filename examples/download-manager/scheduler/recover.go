package scheduler

import (
	"context"
	"sort"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// Recover loads active jobs from the store and enqueues the queued ones.
// Call it once, after store.Reconcile, before Start.
func (e *Engine) Recover(ctx context.Context) error {
	jobs, err := e.opts.Store.ActiveJobs(ctx)
	if err != nil {
		return err
	}

	e.mu.Lock()
	for _, job := range jobs {
		if _, seen := e.jobs[job.ID]; seen {
			continue
		}

		l := &live{job: job}
		e.jobs[job.ID] = l

		if job.State == domain.StateQueued && len(e.queue) < e.opts.Queue {
			l.queued = true
			e.queue = append(e.queue, job.ID)
		}
	}
	e.mu.Unlock()

	e.signal()

	return nil
}

// Job returns one live job.
func (e *Engine) Job(id string) (domain.Job, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	l, ok := e.jobs[id]
	if !ok {
		return domain.Job{}, false
	}

	return l.job, true
}

// Active returns the queued, running, and paused jobs, oldest first.
func (e *Engine) Active() []domain.Job {
	e.mu.Lock()
	defer e.mu.Unlock()

	out := make([]domain.Job, 0, len(e.jobs))

	for _, l := range e.jobs {
		if l.job.State.Active() {
			out = append(out, l.job)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].CreatedAt.Equal(out[j].CreatedAt) {
			return out[i].ID < out[j].ID
		}

		return out[i].CreatedAt.Before(out[j].CreatedAt)
	})

	return out
}

// Drain pops at most max events.
func (e *Engine) Drain(max int) []Event { return e.out.Drain(max) }

// Updated returns a signal channel that fires when events arrive.
func (e *Engine) Updated() <-chan struct{} { return e.out.Updated() }
