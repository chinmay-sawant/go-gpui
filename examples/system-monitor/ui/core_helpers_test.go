package ui

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
)

// waitFor polls cond until it holds or the deadline passes.
func waitFor(t *testing.T, what string, cond func() bool) {
	t.Helper()

	deadline := time.Now().Add(5 * time.Second)

	for !cond() {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s", what)
		}

		time.Sleep(5 * time.Millisecond)
	}
}

// dummyManager starts a fast dummy collector for tests.
func dummyManager(t *testing.T) *collector.Manager {
	t.Helper()

	mgr := collector.New(collector.Options{
		Mode:            collector.ModeDummy,
		Seed:            7,
		SummaryInterval: 10 * time.Millisecond,
		ProcessInterval: 20 * time.Millisecond,
		DetailInterval:  20 * time.Millisecond,
	})
	if err := mgr.Start(context.Background()); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = mgr.Close(context.Background()) })

	return mgr
}
