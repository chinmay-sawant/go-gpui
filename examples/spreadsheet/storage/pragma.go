package storage

import (
	"context"
	"database/sql"
)

// Checkpoint folds the WAL back into the database. Callers run it off the
// UI loop, never between an edit and its acknowledgement.
func (s *Store) Checkpoint(ctx context.Context) error {
	return s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, "PRAGMA wal_checkpoint(TRUNCATE)")

		return err
	})
}

// readJournal asks SQLite which journal mode is active.
func (s *Store) readJournal() string {
	var mode string

	err := s.do(context.Background(), func(ctx context.Context, db *sql.DB) error {
		return db.QueryRowContext(ctx, "PRAGMA journal_mode").Scan(&mode)
	})
	if err != nil {
		return "unknown"
	}

	return mode
}
