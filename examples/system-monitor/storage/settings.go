package storage

import (
	"context"
	"time"
)

// Setting is one saved preference, such as the theme or the selected mode.
type Setting struct {
	Key     string
	Value   string
	Updated time.Time
}

// Settings returns every setting, sorted by key.
func (s *Store) Settings(ctx context.Context) ([]Setting, error) { return nil, errNotImplemented }

// Setting returns one value and whether it exists.
func (s *Store) Setting(ctx context.Context, key string) (string, bool, error) {
	return "", false, errNotImplemented
}

// SetSetting writes one setting, replacing an earlier value.
func (s *Store) SetSetting(ctx context.Context, key, value string) error { return errNotImplemented }

// DeleteSetting removes one setting.
func (s *Store) DeleteSetting(ctx context.Context, key string) error { return errNotImplemented }
