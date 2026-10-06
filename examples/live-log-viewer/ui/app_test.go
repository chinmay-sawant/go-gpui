package ui

import (
	"context"
	"testing"
	"time"
)

// pumpUntil ticks the app until cond is true or the deadline passes.
func pumpUntil(t *testing.T, a *App, cond func() bool, what string) {
	t.Helper()

	deadline := time.Now().Add(10 * time.Second)
	ctx := context.Background()

	for time.Now().Before(deadline) {
		if cond() {
			return
		}

		if err := a.Tick(ctx); err != nil {
			t.Fatal(err)
		}

		time.Sleep(2 * time.Millisecond)
	}

	t.Fatalf("timeout waiting for %s", what)
}

func newApp(t *testing.T, ff *fakeFeed, opts Options) *App {
	t.Helper()

	opts.Feed = ff
	if opts.Poll == 0 {
		opts.Poll = 20 * time.Millisecond
	}

	a, err := New(opts)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = a.Close() })

	return a
}

// lastID returns the newest loaded entry ID.
func (a *App) lastID() int64 {
	if n := a.pager.Len(); n > 0 {
		return a.pager.Entries[n-1].ID
	}

	return 0
}
