package scheduler

import (
	"context"
	"path/filepath"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// Add validates the request, reserves a destination, persists the queued
// job, and enqueues it. Creation is durable before any transfer starts.
func (e *Engine) Add(ctx context.Context, req AddRequest) (domain.Job, error) {
	if err := validateAdd(req); err != nil {
		return domain.Job{}, err
	}

	name := req.Name
	if name == "" {
		name = transfer.SafeName(req.URL)
	}

	now := e.opts.Now()
	job := domain.Job{
		ID:        newID(),
		URL:       req.URL,
		Name:      name,
		State:     domain.StateQueued,
		Expected:  expectedOf(req.Expected),
		Checksum:  req.Checksum,
		CreatedAt: now,
		UpdatedAt: now,
	}

	e.mu.Lock()
	if e.closed {
		e.mu.Unlock()

		return domain.Job{}, ErrClosed
	}

	if len(e.queue) >= e.opts.Queue {
		e.mu.Unlock()

		return domain.Job{}, ErrQueueFull
	}

	dest, err := e.reserveDestinationLocked(req.Dir, name)
	if err != nil {
		e.mu.Unlock()

		return domain.Job{}, err
	}

	job.Destination = dest
	job.Name = filepath.Base(dest)
	e.mu.Unlock()

	if err := e.opts.Store.SaveJob(ctx, job); err != nil {
		return domain.Job{}, err
	}

	e.mu.Lock()
	e.jobs[job.ID] = &live{job: job, queued: true}
	e.queue = append(e.queue, job.ID)
	e.mu.Unlock()

	e.signal()
	e.out.publish(Event{Kind: EventState, Job: job, At: now})

	return job, nil
}
