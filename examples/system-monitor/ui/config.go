package ui

import (
	"context"
	"errors"
)

// ErrNoSource means New was called without a data source.
var ErrNoSource = errors.New("system-monitor/ui: no data source")

// Config names the window, the data source, and the settings store.
type Config struct {
	// Source feeds samples and process snapshots. It must not be nil.
	Source Source
	// Store persists the theme toggle. A nil store keeps it in memory.
	Store Store
	// Live starts on the live source. The caller builds the source for the
	// mode; the mode toggle calls SetLive on it.
	Live bool
	// Width and Height are the first frame, in CSS pixels.
	Width, Height int
	// MinWidth and MinHeight keep the layout usable on small screens.
	MinWidth, MinHeight int
	// noPump stops the poll goroutine. Tests set it; production polls.
	noPump bool
}

// Window defaults. The minimum keeps the overview grid and the process
// table usable, so a small window never needs horizontal scrolling.
const (
	defaultWidth  = 1280
	defaultHeight = 720
	minWidth      = 900
	minHeight     = 600
)

// withDefaults fills zero window sizes.
func withDefaults(cfg Config) Config {
	if cfg.Width <= 0 {
		cfg.Width = defaultWidth
	}

	if cfg.Height <= 0 {
		cfg.Height = defaultHeight
	}

	if cfg.MinWidth <= 0 {
		cfg.MinWidth = minWidth
	}

	if cfg.MinHeight <= 0 {
		cfg.MinHeight = minHeight
	}

	return cfg
}

// Store persists the few settings a user changes.
type Store interface {
	LoadSettings(ctx context.Context) (Settings, error)
	SaveSettings(ctx context.Context, s Settings) error
}

// Settings is the persisted user state.
type Settings struct {
	Dark bool
}
