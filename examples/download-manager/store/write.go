package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Checkpoint writes only Done, Total, and UpdatedAt. It cannot clobber a
// concurrent edit of names, state, or validators.
func (s *Store) Checkpoint(ctx context.Context, id string, done, total int64, at time.Time) error {
	return s.do(ctx, func(ctx context.Context) error {
		res, err := s.db.ExecContext(ctx,
			`UPDATE jobs SET done = ?, total = ?, updated_ms = ? WHERE id = ?`,
			done, total, msOf(at), id)
		if err != nil {
			return err
		}

		return oneRow(res, id)
	})
}

// DeleteJob removes one row.
func (s *Store) DeleteJob(ctx context.Context, id string) error {
	return s.do(ctx, func(ctx context.Context) error {
		res, err := s.db.ExecContext(ctx, `DELETE FROM jobs WHERE id = ?`, id)
		if err != nil {
			return err
		}

		return oneRow(res, id)
	})
}

// oneRow maps zero affected rows to ErrNotFound.
func oneRow(res sql.Result, id string) error {
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("%w: %s", ErrNotFound, id)
	}

	return nil
}
