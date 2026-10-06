package storage

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// Recorder owns one goroutine that writes recorded samples to one session
// over a bounded queue. Record and RecordSample never block: when the queue
// is full the sample is dropped and counted, so a slow disk cannot stall the
// collector. The Recorder satisfies collector.Sink without importing it.
type Recorder struct {
	store     *Store
	sessionID int64
	queue     chan []Row

	mu        sync.Mutex
	written   int64
	failed    int64
	lastError string

	dropped  atomic.Uint64
	closed   atomic.Bool
	done     chan struct{}
	closeOne sync.Once
}

// NewRecorder starts the writer. queue <= 0 takes a default of 64 samples.
func NewRecorder(st *Store, sessionID int64, queue int) *Recorder {
	if queue <= 0 {
		queue = 64
	}

	r := &Recorder{
		store:     st,
		sessionID: sessionID,
		queue:     make(chan []Row, queue),
		done:      make(chan struct{}),
	}

	go r.run()

	return r
}

// Close stops the writer after the queued rows are written or ctx is done. A
// failed write is reported, never hidden.
func (r *Recorder) Close(ctx context.Context) error {
	r.closeOne.Do(func() {
		r.closed.Store(true)
		close(r.queue)
	})

	select {
	case <-r.done:
		stats := r.Stats()
		if stats.Failed > 0 {
			return fmt.Errorf("storage: recorder wrote %d samples, %d failed: %s",
				stats.Written, stats.Failed, stats.LastError)
		}

		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
