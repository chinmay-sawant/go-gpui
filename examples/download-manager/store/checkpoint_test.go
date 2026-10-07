package store

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestCheckpointTouchesOnlyProgress leaves other columns alone.
func TestCheckpointTouchesOnlyProgress(t *testing.T) {
	s := newStore(t)
	ctx := context.Background()

	job := testJob("cp")
	job.Expected = 1000

	if err := s.SaveJob(ctx, job); err != nil {
		t.Fatal(err)
	}

	at := job.UpdatedAt.Add(time.Minute)
	if err := s.Checkpoint(ctx, "cp", 500, 1000, at); err != nil {
		t.Fatal(err)
	}

	got, err := s.Job(ctx, "cp")
	if err != nil {
		t.Fatal(err)
	}

	if got.Done != 500 || got.State != domain.StateQueued || got.URL != job.URL {
		t.Errorf("checkpoint clobbered a column: %+v", got)
	}

	if err := s.Checkpoint(ctx, "missing", 1, 1, at); err == nil {
		t.Error("checkpoint on a missing job succeeded")
	}
}
