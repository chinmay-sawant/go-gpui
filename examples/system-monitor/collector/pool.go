package collector

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// Pool runs source calls on a fixed set of workers over a bounded queue.
// TryDo never blocks: when the queue is full it returns false and the caller
// skips that tick, which keeps delivery at the latest snapshot instead of a
// backlog. Each call gets a context deadline; a call that overruns it is
// abandoned and its worker returns to service when the call finally ends.
type Pool struct {
	jobs    chan job
	timeout time.Duration
	workers int

	mu     sync.Mutex
	closed bool
	wg     sync.WaitGroup

	busy      atomic.Int64
	accepted  atomic.Uint64
	dropped   atomic.Uint64
	completed atomic.Uint64
	timedOut  atomic.Uint64
}

type job struct {
	ctx    context.Context
	fn     func(context.Context) error
	onDone func(error)
}

// NewPool starts workers goroutines that pull from a queue of size queue.
func NewPool(workers, queue int, timeout time.Duration) *Pool {
	if workers < 1 {
		workers = 1
	}
	if queue < 1 {
		queue = 1
	}
	if timeout <= 0 {
		timeout = time.Second
	}

	p := &Pool{jobs: make(chan job, queue), timeout: timeout, workers: workers}
	p.wg.Add(workers)

	for range workers {
		go p.worker()
	}

	return p
}

// TryDo queues fn when a queue slot is free, and reports whether it did.
// onDone may be nil and runs after fn, with fn's error, on the worker.
func (p *Pool) TryDo(ctx context.Context, fn func(context.Context) error, onDone func(error)) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return false
	}

	select {
	case p.jobs <- job{ctx: ctx, fn: fn, onDone: onDone}:
		p.accepted.Add(1)

		return true
	default:
		p.dropped.Add(1)

		return false
	}
}

// Close stops the pool after the calls in flight finish.
func (p *Pool) Close() {
	p.mu.Lock()
	if !p.closed {
		p.closed = true
		close(p.jobs)
	}
	p.mu.Unlock()

	p.wg.Wait()
}
