// Package scheduler owns the bounded worker pool: it queues jobs, starts
// downloads through a transfer.Transport, persists transitions, and
// publishes coalesced progress plus undroppable state events. All exported
// methods are safe to call from any goroutine.
package scheduler

import (
	"context"
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// ErrClosed is returned after Close.
var ErrClosed = errors.New("scheduler: engine is closed")

// Store is the persistence seam. store.Store satisfies it.
type Store interface {
	SaveJob(ctx context.Context, j domain.Job) error
	Checkpoint(ctx context.Context, id string, done, total int64, at time.Time) error
	Job(ctx context.Context, id string) (domain.Job, error)
	ActiveJobs(ctx context.Context) ([]domain.Job, error)
}

// Options configures an Engine.
type Options struct {
	// Transport runs one job. Required.
	Transport transfer.Transport
	// Store persists job rows. Required.
	Store Store
	// Workers is the pool size. Zero means three.
	Workers int
	// Queue bounds the waiting list. Zero means 64.
	Queue int
	// ProgressEvery bounds how often one job's bytes are persisted.
	// Zero means 500 ms.
	ProgressEvery time.Duration
	// ShutdownBudget bounds Close. Zero means five seconds.
	ShutdownBudget time.Duration
	// Now replaces the clock in tests. Nil means time.Now.
	Now func() time.Time
}

// AddRequest asks for one new job.
type AddRequest struct {
	URL      string
	Dir      string
	Name     string
	Expected int64
	Checksum string
}
