package store

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestSummaryAndCleanup checks counts and bounded deletion.
func TestSummaryAndCleanup(t *testing.T) {
	s := newStore(t)
	ctx := context.Background()
	base := time.Now().UTC().Truncate(time.Millisecond)

	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("c-%03d", i)
		if err := s.SaveJob(ctx, terminalJob(id, base.Add(time.Duration(i)*time.Second))); err != nil {
			t.Fatal(err)
		}
	}

	if err := s.SaveJob(ctx, testJob("active-1")); err != nil {
		t.Fatal(err)
	}

	counts, err := s.Summary(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if counts.Total != 11 || counts.ByState[domain.StateCompleted] != 10 ||
		counts.ByState[domain.StateQueued] != 1 {
		t.Fatalf("summary %+v", counts)
	}

	removed, err := s.Cleanup(ctx, 4, 3)
	if err != nil {
		t.Fatal(err)
	}

	if removed != 3 {
		t.Fatalf("first batch removed %d", removed)
	}

	total := 0
	for {
		n, err := s.Cleanup(ctx, 4, 3)
		if err != nil {
			t.Fatal(err)
		}

		if n == 0 {
			break
		}

		total += n
	}

	if total != 3 {
		t.Fatalf("later batches removed %d, want 3", total)
	}

	after, _ := s.Summary(ctx)
	if after.Total != 5 || after.ByState[domain.StateQueued] != 1 {
		t.Fatalf("after cleanup %+v", after)
	}
}
