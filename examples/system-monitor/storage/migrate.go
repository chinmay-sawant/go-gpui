package storage

import (
	"context"
	"fmt"
)

// migration is one schema step. The version is the user_version it moves to.
type migration struct {
	version int
	sql     string
}

// migrations run in order inside one transaction.
var migrations = []migration{
	{version: 1, sql: schemaSQL},
}

// migrate brings the file to SchemaVersion. A file from a newer build is
// refused before anything is written, so it stays untouched.
func (s *Store) migrate(ctx context.Context) error {
	var version int

	if err := s.db.QueryRowContext(ctx, "PRAGMA user_version").Scan(&version); err != nil {
		return err
	}

	if version > SchemaVersion {
		return fmt.Errorf("%w (file has %d, build has %d)", ErrSchemaNewer, version, SchemaVersion)
	}
	if version == SchemaVersion {
		return nil
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	for _, m := range migrations {
		if m.version <= version {
			continue
		}
		if _, err := tx.ExecContext(ctx, m.sql); err != nil {
			return err
		}
	}

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("PRAGMA user_version = %d", SchemaVersion)); err != nil {
		return err
	}

	return tx.Commit()
}

// SchemaVersionOf returns the file's recorded schema version.
func (s *Store) SchemaVersionOf(ctx context.Context) (int, error) {
	if err := s.ready(); err != nil {
		return 0, err
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	var version int
	if err := s.db.QueryRowContext(ctx, "PRAGMA user_version").Scan(&version); err != nil {
		return 0, err
	}

	return version, nil
}
