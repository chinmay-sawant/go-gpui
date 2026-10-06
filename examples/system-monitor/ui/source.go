package ui

import (
	"context"
	"time"
)

// Source is the data feed the app polls. Every method must be safe to call
// from a goroutine and must not block for long; the app polls from its own
// goroutines, never from a handler or a tick. The live collector, the dummy
// fixture, and test fakes all satisfy it.
type Source interface {
	// Summary returns the newest system sample. ok is false before the
	// first sample of the current mode.
	Summary(ctx context.Context) (Summary, bool)
	// Processes returns the newest published process table.
	Processes(ctx context.Context) (ProcSnapshot, bool)
	// Track selects the identity the detail view follows; empty clears.
	Track(id string)
	// Tracked returns the state of the tracked process.
	Tracked(ctx context.Context) Tracked
	// Problems returns the latest collector failures as short lines.
	Problems(ctx context.Context) []string
	// SetLive switches the source mode. The source resets its baselines.
	SetLive(live bool) error
	// Close stops the source and releases its workers.
	Close() error
}

// Summary is one system sample: CPU, memory, disk, and network readings.
type Summary struct {
	At       time.Time
	Readings []Reading
}

// Reading is one overview metric. ID is cpu, mem, disk, or net. Scale is the
// full-height value for the graph; zero means the graph scales to its own
// peak. Text is the primary label and Sub the secondary line. OK=false
// paints as unavailable instead of a misleading zero.
type Reading struct {
	ID    string
	Value float64
	OK    bool
	Scale float64
	Text  string
	Sub   string
}
