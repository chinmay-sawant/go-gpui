package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

// fixtureKey stores the version of the seeded dummy fixture.
const fixtureKey = "fixture.dummy.version"

// FixtureVersion returns the stored fixture version, or zero when none was
// seeded.
func (s *Store) FixtureVersion(ctx context.Context) (int, error) {
	if err := s.ready(); err != nil {
		return 0, err
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	var raw string

	err := s.db.QueryRowContext(ctx, `SELECT value FROM meta WHERE key = ?`, fixtureKey).Scan(&raw)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	version, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("storage: fixture version %q: %w", raw, err)
	}

	return version, nil
}

// nameOr returns name or fallback when name is empty.
func nameOr(name, fallback string) string {
	if name == "" {
		return fallback
	}

	return name
}

// markFixtureTx records the fixture version inside a seeding transaction.
func markFixtureTx(ctx context.Context, tx *sql.Tx, version int) error {
	_, err := tx.ExecContext(ctx, `
INSERT INTO meta(key, value) VALUES(?, ?)
ON CONFLICT(key) DO UPDATE SET value = excluded.value`,
		fixtureKey, strconv.Itoa(version))

	return err
}
