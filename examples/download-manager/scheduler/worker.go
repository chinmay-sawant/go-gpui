package scheduler

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// workerLoop takes jobs until the pool context ends.
func (e *Engine) workerLoop(ctx context.Context) {
	defer e.wg.Done()

	for {
		l, ok := e.next(ctx)
		if !ok {
			return
		}

		e.run(ctx, l)
	}
}

// next pops the oldest queued job, waiting for work or the context.
func (e *Engine) next(ctx context.Context) (*live, bool) {
	for {
		e.mu.Lock()
		if e.closed {
			e.mu.Unlock()

			return nil, false
		}

		if len(e.queue) > 0 {
			id := e.queue[0]
			e.queue = e.queue[1:]
			l := e.jobs[id]

			if l != nil && l.queued && !l.running && l.job.State == domain.StateQueued {
				l.queued = false
				l.running = true
				e.mu.Unlock()

				return l, true
			}

			e.mu.Unlock()

			continue
		}

		e.mu.Unlock()

		select {
		case <-ctx.Done():
			return nil, false
		case <-e.wake:
		}
	}
}
