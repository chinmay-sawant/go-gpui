package ui

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/storage"
)

// themeKey is the storage setting that carries the theme toggle.
const themeKey = "ui.theme"

// CollectorStore adapts the example's storage to the Store seam. A read
// falls back to defaults when storage is unavailable; a failed write is
// reported, never silently treated as saved.
type CollectorStore struct {
	st *storage.Store
}

// NewCollectorStore wraps an open store.
func NewCollectorStore(st *storage.Store) *CollectorStore {
	return &CollectorStore{st: st}
}

// LoadSettings reads the saved theme.
func (s *CollectorStore) LoadSettings(ctx context.Context) (Settings, error) {
	v, ok, err := s.st.Setting(ctx, themeKey)
	if err != nil || !ok {
		return Settings{}, err
	}

	return Settings{Dark: v == "dark"}, nil
}

// SaveSettings writes the theme.
func (s *CollectorStore) SaveSettings(ctx context.Context, v Settings) error {
	theme := "light"
	if v.Dark {
		theme = "dark"
	}

	return s.st.SetSetting(ctx, themeKey, theme)
}
