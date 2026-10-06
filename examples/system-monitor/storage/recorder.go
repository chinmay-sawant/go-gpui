package storage

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
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
	return nil
}

// RecordSample converts one sample to rows and enqueues them. It reports
// whether the sample was queued.
func (r *Recorder) RecordSample(s domain.Sample) bool { return false }

// Record enqueues prepared rows.
func (r *Recorder) Record(rows []Row) bool { return false }

// RecorderStats reports what the recorder wrote, dropped, and failed.
type RecorderStats struct {
	Written   int64
	Failed    int64
	Dropped   uint64
	Queued    int
	LastError string
}

// Stats returns a point-in-time view.
func (r *Recorder) Stats() RecorderStats { return RecorderStats{} }

// Close stops the writer after the queued rows are written or ctx is done.
// A failed final write is reported in an error.
func (r *Recorder) Close(ctx context.Context) error { return errNotImplemented }
