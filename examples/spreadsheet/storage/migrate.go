package storage

import (
	"context"
	"database/sql"
	"fmt"
)

// SchemaVersion is the schema this build writes.
const SchemaVersion = 1

type migration struct {
	version int
	name    string
	sql     string
}

// migrations apply in order; each runs in its own transaction.
var migrations = []migration{{version: 1, name: "initial", sql: schemaV1}}

// migrate brings the database up to SchemaVersion.
func (s *Store) migrate(ctx context.Context) error {
	return s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		return applyMigrations(ctx, db, SchemaVersion, migrations)
	})
}

// applyMigrations is the testable core of migrate. A database from a newer
// build is rejected before anything writes.
func applyMigrations(ctx context.Context, db *sql.DB, newest int, migs []migration) error {
	var version int

	if err := db.QueryRowContext(ctx, `PRAGMA user_version`).Scan(&version); err != nil {
		return err
	}

	if version > newest {
		return fmt.Errorf("%w: database %d, build %d", ErrSchemaNewer, version, newest)
	}

	for _, m := range migs {
		if m.version <= version {
			continue
		}

		if m.version > newest {
			return fmt.Errorf("%w: migration %d", ErrSchemaNewer, m.version)
		}

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		if _, err := tx.ExecContext(ctx, m.sql); err != nil {
			return rollback(tx, fmt.Errorf("migration %s: %w", m.name, err))
		}

		if _, err := tx.ExecContext(ctx, fmt.Sprintf("PRAGMA user_version = %d", m.version)); err != nil {
			return rollback(tx, fmt.Errorf("migration %s: %w", m.name, err))
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("migration %s: %w", m.name, err)
		}
	}

	return nil
}
