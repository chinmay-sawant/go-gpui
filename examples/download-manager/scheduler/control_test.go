package scheduler

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestPauseAndResumeRunning stops mid-flight and finishes after resume.
func TestPauseAndResumeRunning(t *testing.T) {
	gate := newGate()
	eng, _ := newTestEngine(t, gate, Options{Workers: 1})
	eng.Start(context.Background())

	job := add(t, eng, t.TempDir(), "https://example.invalid/pause.bin")
	waitStarted(t, gate, job.ID)

	if err := eng.Pause(context.Background(), job.ID); err != nil {
		t.Fatal(err)
	}

	paused := waitState(t, eng, job.ID, domain.StatePaused)
	if paused.Done != 5 {
		t.Errorf("paused at %d bytes, want the last report", paused.Done)
	}

	if err := eng.Resume(context.Background(), job.ID); err != nil {
		t.Fatal(err)
	}

	gate.open()
	waitState(t, eng, job.ID, domain.StateCompleted)
}
