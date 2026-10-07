package store

import (
	"context"
	"os"
	"path/filepath"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// reconcilePartial repairs a queued or paused row against its partial file
// and renames a destination an outside process reused.
func (s *Store) reconcilePartial(ctx context.Context, job domain.Job, rep *Report) error {
	size, err := optionalSize(transfer.PartialPath(job.Destination))
	if err != nil {
		return err
	}

	original := job.Destination

	if err := s.renameReused(&job, rep); err != nil {
		return err
	}

	if job.Destination != original {
		if err := s.setDestination(ctx, job.ID, job.Destination); err != nil {
			return err
		}
	}

	if job.Expected > 0 && size > job.Expected {
		if err := os.Truncate(transfer.PartialPath(job.Destination), 0); err != nil {
			return err
		}

		rep.Reset++

		return s.mark(ctx, job, job.State, 0, "")
	}

	return s.mark(ctx, job, job.State, size, "")
}

// renameReused gives an active job a new destination when something else
// took its path and the partial does not prove ownership.
func (s *Store) renameReused(job *domain.Job, rep *Report) error {
	size, err := fileSize(job.Destination)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	if adoptable(*job, size) {
		return nil
	}

	fresh, err := transfer.UniqueDestination(filepath.Dir(job.Destination), job.Name)
	if err != nil {
		return err
	}

	job.Destination = fresh
	rep.Renamed++

	return nil
}
