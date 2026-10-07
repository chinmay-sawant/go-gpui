package scheduler

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestCancelRunning marks a live transfer cancelled.
func TestCancelRunning(t *testing.T) {
	gate := newGate()
	eng, st := newTestEngine(t, gate, Options{Workers: 1})
	eng.Start(context.Background())

	job := add(t, eng, t.TempDir(), "https://example.invalid/cancel.bin")
	waitStarted(t, gate, job.ID)

	if err := eng.Cancel(context.Background(), job.ID); err != nil {
		t.Fatal(err)
	}

	waitState(t, eng, job.ID, domain.StateCancelled)

	row, err := st.Job(context.Background(), job.ID)
	if err != nil {
		t.Fatal(err)
	}

	if row.State != domain.StateCancelled {
		t.Errorf("persisted state %s", row.State)
	}
}
