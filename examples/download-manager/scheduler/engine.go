package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// live is the in-memory mirror of one job.
type live struct {
	job     domain.Job
	cancel  context.CancelFunc
	queued  bool
	running bool
	err     error // last transfer error, set by the worker
}

// Engine is the bounded worker pool.
type Engine struct {
	opts  Options
	out   *Outbox
	mu    sync.Mutex
	jobs  map[string]*live
	queue []string
	wake  chan struct{}
	wg    sync.WaitGroup

	started bool
	closed  bool
}

// New builds an engine. Transport and Store are required.
func New(opts Options) (*Engine, error) {
	if opts.Workers <= 0 {
		opts.Workers = 3
	}

	if opts.Queue <= 0 {
		opts.Queue = 64
	}

	if opts.Now == nil {
		opts.Now = timeNow
	}

	if opts.ProgressEvery <= 0 {
		opts.ProgressEvery = 500 * time.Millisecond
	}

	if opts.ShutdownBudget <= 0 {
		opts.ShutdownBudget = 5 * time.Second
	}

	if opts.Transport == nil || opts.Store == nil {
		return nil, errMissingOption
	}

	return &Engine{
		opts: opts,
		out:  NewOutbox(),
		jobs: map[string]*live{},
		wake: make(chan struct{}, 1),
	}, nil
}
