package store

import (
	"context"
	"os"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// reconcileRunning decides between adopting a finalized file, finalizing a
// complete partial, or falling back to paused.
func (s *Store) reconcileRunning(ctx context.Context, job domain.Job, rep *Report) error {
	size, err := fileSize(job.Destination)

	switch {
	case err == nil && adoptable(job, size):
		rep.Completed++

		return s.mark(ctx, job, domain.StateCompleted, size, "")
	case err != nil && !os.IsNotExist(err):
		return err
	}

	partialPath := transfer.PartialPath(job.Destination)

	partial, err := optionalSize(partialPath)
	if err != nil {
		return err
	}

	if job.Expected > 0 && partial > job.Expected {
		if err := os.Truncate(partialPath, 0); err != nil {
			return err
		}

		rep.Reset++

		partial = 0
	}

	if job.Expected > 0 && partial == job.Expected {
		if err := transfer.CheckChecksum(partialPath, job.Checksum); err == nil {
			err := transfer.Finalize(partialPath, job.Destination,
				transfer.FinalizeOptions{})
			if err == nil {
				rep.Completed++

				return s.mark(ctx, job, domain.StateCompleted, partial, "")
			}
		}
	}

	rep.Recovered++

	return s.mark(ctx, job, domain.StatePaused, partial, "interrupted; resume or retry")
}

// reconcileFinished fails a completed row whose file is gone.
func (s *Store) reconcileFinished(ctx context.Context, job domain.Job, rep *Report) error {
	_, err := fileSize(job.Destination)
	if err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	rep.Missing++

	return s.mark(ctx, job, domain.StateFailed, job.Done, "destination file is missing")
}
