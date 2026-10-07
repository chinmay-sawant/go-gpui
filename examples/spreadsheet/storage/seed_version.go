package storage

import (
	"context"
	"database/sql"
	"errors"
)

// SeedVersion reads the marker; zero means the database was never seeded.
func (s *Store) SeedVersion(ctx context.Context) (int, error) {
	var version int

	err := s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		err := db.QueryRowContext(ctx,
			`SELECT CAST(value AS INTEGER) FROM meta WHERE key = 'seed_version'`).Scan(&version)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})

	return version, err
}
