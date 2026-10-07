package store

import (
	"context"
)

// Seed markers. A later build appends a new version instead of editing
// this one; INSERT OR IGNORE leaves user edits alone.
const (
	seedVersionKey   = "dummy_seed_version"
	dummySeedVersion = "1"
)

// SeedDummy inserts jobCount deterministic rows once per seed version.
// Reopening never duplicates them and never rewrites user changes.
func (s *Store) SeedDummy(ctx context.Context, jobCount int) error {
	if jobCount <= 0 {
		return nil
	}

	return s.do(ctx, func(ctx context.Context) error {
		return s.seedLocked(ctx, jobCount)
	})
}

// seedLocked checks the marker and inserts inside one transaction.
func (s *Store) seedLocked(ctx context.Context, jobCount int) error {
	version, found, err := s.meta(ctx, seedVersionKey)
	if err != nil {
		return err
	}

	if found && version == dummySeedVersion {
		return nil
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() { _ = tx.Rollback() }()

	now := timeNow()

	for i := 0; i < jobCount; i++ {
		job := dummyJob(i, jobCount, now)
		if _, err := tx.ExecContext(ctx,
			`INSERT OR IGNORE INTO jobs (`+jobColumns+
				`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			jobArgs(job)...); err != nil {
			return err
		}
	}

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO meta (key, value) VALUES (?, ?)
		 ON CONFLICT(key) DO UPDATE SET value = excluded.value`,
		seedVersionKey, dummySeedVersion); err != nil {
		return err
	}

	return tx.Commit()
}
