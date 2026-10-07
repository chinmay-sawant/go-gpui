package store

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// upsertJob inserts or replaces a whole row.
const upsertJob = `INSERT INTO jobs (` + jobColumns + `)
VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
ON CONFLICT(id) DO UPDATE SET
	url = excluded.url, destination = excluded.destination,
	name = excluded.name, state = excluded.state, done = excluded.done,
	total = excluded.total, expected = excluded.expected,
	etag = excluded.etag, last_modified = excluded.last_modified,
	checksum = excluded.checksum, attempts = excluded.attempts,
	error = excluded.error, created_ms = excluded.created_ms,
	updated_ms = excluded.updated_ms`

// SaveJob writes a whole job row. The caller owns state validation; the
// store writes what it is given and returns constraint errors unchanged.
func (s *Store) SaveJob(ctx context.Context, j domain.Job) error {
	if !j.State.Valid() {
		return fmt.Errorf("store: invalid state %q", j.State)
	}

	if j.UpdatedAt.IsZero() {
		j.UpdatedAt = timeNow()
	}

	if j.CreatedAt.IsZero() {
		j.CreatedAt = j.UpdatedAt
	}

	return s.do(ctx, func(ctx context.Context) error {
		_, err := s.db.ExecContext(ctx, upsertJob, jobArgs(j)...)

		return err
	})
}
