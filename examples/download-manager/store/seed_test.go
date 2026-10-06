package store

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestSeedDummyIsIdempotent seeds 100 rows once and never again.
func TestSeedDummyIsIdempotent(t *testing.T) {
	s := newStore(t)
	ctx := context.Background()

	if err := s.SeedDummy(ctx, 100); err != nil {
		t.Fatal(err)
	}

	counts, err := s.Summary(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if counts.Total != 100 {
		t.Fatalf("seeded %d rows, want 100", counts.Total)
	}

	for _, state := range []domain.State{domain.StateCompleted, domain.StateFailed, domain.StateQueued, domain.StatePaused} {
		if counts.ByState[state] == 0 {
			t.Errorf("no seeded rows in %s", state)
		}
	}

	if err := s.SeedDummy(ctx, 100); err != nil {
		t.Fatal(err)
	}

	again, _ := s.Summary(ctx)
	if again.Total != 100 {
		t.Errorf("second seed duplicated rows: %d", again.Total)
	}
}

// TestSeedPreservesUserEdits keeps a renamed job through a reseed.
func TestSeedPreservesUserEdits(t *testing.T) {
	s := newStore(t)
	ctx := context.Background()

	if err := s.SeedDummy(ctx, 20); err != nil {
		t.Fatal(err)
	}

	job, err := s.Job(ctx, "dummy-0000")
	if err != nil {
		t.Fatal(err)
	}

	job.Name = "renamed-by-user"
	job.URL = "https://user.example/renamed"
	job.State = domain.StateCompleted
	job.Done = 1

	if err := s.SaveJob(ctx, job); err != nil {
		t.Fatal(err)
	}

	if err := s.SeedDummy(ctx, 20); err != nil {
		t.Fatal(err)
	}

	got, err := s.Job(ctx, "dummy-0000")
	if err != nil {
		t.Fatal(err)
	}

	if got.Name != "renamed-by-user" || got.URL != "https://user.example/renamed" {
		t.Errorf("reseed overwrote user edits: %+v", got)
	}
}
