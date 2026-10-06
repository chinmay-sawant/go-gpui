package store

import (
	"context"
	"strings"
)

// enableJournal prefers WAL for local files and keeps the rollback journal
// when the storage refuses WAL, for example on a network share. WAL still
// allows a single writer; every statement here goes through the one
// connection.
func (s *Store) enableJournal(ctx context.Context) error {
	if s.memory {
		s.journal = "memory"

		return nil
	}

	var mode string

	err := s.db.QueryRowContext(ctx, `PRAGMA journal_mode=WAL`).Scan(&mode)
	if err == nil && strings.EqualFold(mode, "wal") {
		s.journal = "wal"
		_, err := s.db.ExecContext(ctx, `PRAGMA synchronous=NORMAL`)

		return err
	}

	s.journal = "delete"
	_, err = s.db.ExecContext(ctx, `PRAGMA journal_mode=DELETE`)

	return err
}

// nowMS is the storage clock.
func nowMS() int64 { return msOf(timeNow()) }
