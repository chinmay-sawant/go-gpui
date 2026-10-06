package ui

import (
	"context"
	"sync"
	"time"
)

// fakeSource is an in-memory Source for headless tests.
type fakeSource struct {
	mu       sync.Mutex
	summary  Summary
	haveSum  bool
	procs    ProcSnapshot
	havePr   bool
	tracked  Tracked
	probs    []string
	live     bool
	setCalls int
	closed   bool
	tracks   []string
}

// Summary returns the configured sample.
func (f *fakeSource) Summary(context.Context) (Summary, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.summary, f.haveSum
}

// Processes returns the configured snapshot.
func (f *fakeSource) Processes(context.Context) (ProcSnapshot, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.procs, f.havePr
}

// Track records the identities the app selected.
func (f *fakeSource) Track(id string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.tracks = append(f.tracks, id)
}

// Tracked returns the configured tracked state.
func (f *fakeSource) Tracked(context.Context) Tracked {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.tracked
}

// Problems returns the configured collector problems.
func (f *fakeSource) Problems(context.Context) []string {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.probs
}

// SetLive flips the mode and records the switch.
func (f *fakeSource) SetLive(live bool) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.live = live
	f.setCalls++

	return nil
}

// Close marks the source closed.
func (f *fakeSource) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.closed = true

	return nil
}

// liveMode reports the last requested mode.
func (f *fakeSource) liveMode() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.live
}

// sample builds a one-reading summary.
func sample(id string, v float64) Summary {
	return Summary{
		At: time.Unix(1700000000, 0),
		Readings: []Reading{{
			ID: id, Value: v, OK: true, Scale: 1,
			Text: cpuText(v*100, true), Sub: "test",
		}},
	}
}
