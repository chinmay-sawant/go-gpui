package collector

import "context"

// Do queues fn and waits until a worker has run it or ctx is done. The
// sampling loops use TryDo instead, so a slow source never builds a backlog.
func (p *Pool) Do(ctx context.Context, fn func(context.Context) error) error {
	done := make(chan error, 1)
	if !p.TryDo(ctx, fn, func(err error) { done <- err }) {
		return ErrPoolFull
	}

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PoolStats is a point-in-time view used for diagnostics.
type PoolStats struct {
	Workers   int
	Busy      int
	Queued    int
	Accepted  uint64
	Dropped   uint64
	Completed uint64
	TimedOut  uint64
}

// Stats returns pool counters. TimedOut counts calls that returned the
// deadline error.
func (p *Pool) Stats() PoolStats {
	return PoolStats{
		Workers:   p.workers,
		Busy:      int(p.busy.Load()),
		Queued:    len(p.jobs),
		Accepted:  p.accepted.Load(),
		Dropped:   p.dropped.Load(),
		Completed: p.completed.Load(),
		TimedOut:  p.timedOut.Load(),
	}
}
