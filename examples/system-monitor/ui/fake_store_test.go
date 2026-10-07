package ui

import (
	"context"
	"sync"
)

// fakeStore records the persisted settings.
type fakeStore struct {
	mu    sync.Mutex
	saved Settings
	saves int
	err   error
}

// LoadSettings returns the stored settings.
func (s *fakeStore) LoadSettings(context.Context) (Settings, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.saved, s.err
}

// SaveSettings records the settings.
func (s *fakeStore) SaveSettings(_ context.Context, v Settings) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.saved = v
	s.saves++

	return s.err
}
