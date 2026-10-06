package storage

import (
	"context"
	"database/sql"
	"errors"
)

// GetPref reads one preference. The bool is false when the key is unset.
func (s *Store) GetPref(ctx context.Context, key string) (string, bool, error) {
	var value string

	found := false

	err := s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		err := db.QueryRowContext(ctx, `SELECT value FROM prefs WHERE key = ?`, key).Scan(&value)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		if err != nil {
			return err
		}

		found = true

		return nil
	})

	return value, found, err
}

// SetPref writes one preference.
func (s *Store) SetPref(ctx context.Context, key, value string) error {
	return s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx,
			`INSERT INTO prefs (key, value) VALUES (?, ?)
			 ON CONFLICT (key) DO UPDATE SET value = excluded.value`, key, value)

		return err
	})
}
