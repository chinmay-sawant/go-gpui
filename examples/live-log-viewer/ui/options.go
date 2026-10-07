package ui

import (
	"errors"
	"time"
)

// Tunables. The README documents each limit.
const (
	// ShutdownBudget bounds how long Close waits for the worker.
	ShutdownBudget = 2 * time.Second
	// PollEvery is the tail poll interval.
	PollEvery = 700 * time.Millisecond
	// SearchDelay is the text search debounce.
	SearchDelay = 250 * time.Millisecond
	// DrainBudget is the most worker results one tick applies.
	DrainBudget = 8
	// DrainTime is the wall budget one tick spends draining.
	DrainTime = 2 * time.Millisecond
	// ExportMax bounds one export selection.
	ExportMax = 5000
	// TailLimit bounds one tail poll fetch.
	TailLimit = 64
)

// ErrShutdown means the worker did not stop inside ShutdownBudget.
var ErrShutdown = errors.New("ui: worker did not stop in the shutdown budget")

// Options configures New.
type Options struct {
	// Feed is the data source; nil renders the shell with no data.
	Feed Feed
	// Dark starts in the dark theme.
	Dark bool
	// ExportDir receives exported CSV files.
	ExportDir string
	// Perf measures redraw and tick time for the status bar.
	Perf bool
	// Poll overrides PollEvery; tests shorten it.
	Poll time.Duration
}
