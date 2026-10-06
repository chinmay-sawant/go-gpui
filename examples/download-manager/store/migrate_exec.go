package store

import (
	"context"
	"database/sql"
	"strings"
)

// apply runs every migration above from, each in its own transaction.
func (s *Store) apply(ctx context.Context, from int) error {
	if _, err := s.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
		version INTEGER PRIMARY KEY,
		applied_ms INTEGER NOT NULL
	)`); err != nil {
		return err
	}

	for _, m := range migrations {
		if m.version <= from {
			continue
		}

		if err := s.applyOne(ctx, m); err != nil {
			return err
		}
	}

	return nil
}

// applyOne commits one migration with its version row or changes nothing.
func (s *Store) applyOne(ctx context.Context, m migration) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := execScript(ctx, tx, m.script); err != nil {
		_ = tx.Rollback()

		return err
	}

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO schema_migrations (version, applied_ms) VALUES (?, ?)`,
		m.version, nowMS()); err != nil {
		_ = tx.Rollback()

		return err
	}

	return tx.Commit()
}

// execScript runs one statement per non-empty, non-comment line.
func execScript(ctx context.Context, tx *sql.Tx, script string) error {
	for _, line := range strings.Split(script, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}

		if _, err := tx.ExecContext(ctx, line); err != nil {
			return err
		}
	}

	return nil
}
