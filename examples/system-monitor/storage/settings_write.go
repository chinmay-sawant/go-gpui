package storage

import (
	"context"
	"errors"
	"time"
)

// SetSetting writes one setting, replacing an earlier value.
func (s *Store) SetSetting(ctx context.Context, key, value string) error {
	if err := s.ready(); err != nil {
		return err
	}
	if key == "" {
		return errors.New("storage: empty setting key")
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	_, err := s.db.ExecContext(ctx, `
INSERT INTO settings(key, value, updated_ns) VALUES(?, ?, ?)
ON CONFLICT(key) DO UPDATE SET value = excluded.value, updated_ns = excluded.updated_ns`,
		key, value, time.Now().UnixNano())

	return err
}

// DeleteSetting removes one setting.
func (s *Store) DeleteSetting(ctx context.Context, key string) error {
	if err := s.ready(); err != nil {
		return err
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	_, err := s.db.ExecContext(ctx, `DELETE FROM settings WHERE key = ?`, key)

	return err
}
