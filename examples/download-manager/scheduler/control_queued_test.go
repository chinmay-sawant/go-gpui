package scheduler

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestPauseResumeCancelQueued drives a job that never started.
func TestPauseResumeCancelQueued(t *testing.T) {
	eng, _ := newTestEngine(t, instantTransport{}, Options{})
	ctx := context.Background()

	job := add(t, eng, t.TempDir(), "https://example.invalid/queued.bin")

	if err := eng.Pause(ctx, job.ID); err != nil {
		t.Fatal(err)
	}

	if got, _ := eng.Job(job.ID); got.State != domain.StatePaused {
		t.Fatalf("state %s", got.State)
	}

	if err := eng.Resume(ctx, job.ID); err != nil {
		t.Fatal(err)
	}

	if got, _ := eng.Job(job.ID); got.State != domain.StateQueued {
		t.Fatalf("state %s", got.State)
	}

	if err := eng.Cancel(ctx, job.ID); err != nil {
		t.Fatal(err)
	}

	if got, _ := eng.Job(job.ID); got.State != domain.StateCancelled {
		t.Fatalf("state %s", got.State)
	}

	if len(eng.Active()) != 0 {
		t.Errorf("active list %+v", eng.Active())
	}
}
