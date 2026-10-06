package store

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// adoptable reports whether a file at Destination can count as the job's
// finished body.
func adoptable(job domain.Job, size int64) bool {
	if job.Expected > 0 {
		if size != job.Expected {
			return false
		}

		return transfer.CheckChecksum(job.Destination, job.Checksum) == nil
	}

	return size > 0
}

// mark writes a reconciled state without touching other columns.
func (s *Store) mark(ctx context.Context, job domain.Job, state domain.State, done int64, note string) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE jobs SET state = ?, done = ?, error = ?, updated_ms = ? WHERE id = ?`,
		string(state), done, note, msOf(timeNow()), job.ID)

	return err
}

// setDestination records a reconciled rename.
func (s *Store) setDestination(ctx context.Context, id, dest string) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE jobs SET destination = ?, updated_ms = ? WHERE id = ?`,
		dest, msOf(timeNow()), id)

	return err
}
