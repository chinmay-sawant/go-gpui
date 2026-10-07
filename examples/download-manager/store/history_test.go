package store

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// terminalJob builds a completed row with a controlled timestamp.
func terminalJob(id string, at time.Time) domain.Job {
	job := testJob(id)
	job.State = domain.StateCompleted
	job.Done = 10
	job.Expected = 10
	job.Total = 10
	job.CreatedAt = at
	job.UpdatedAt = at

	return job
}

// TestHistoryKeysetPaging walks 120 rows in 50-row pages.
func TestHistoryKeysetPaging(t *testing.T) {
	s := newStore(t)
	ctx := context.Background()
	base := time.Now().UTC().Truncate(time.Millisecond)

	for i := 0; i < 120; i++ {
		id := fmt.Sprintf("h-%03d", i)
		if err := s.SaveJob(ctx, terminalJob(id, base.Add(time.Duration(i)*time.Second))); err != nil {
			t.Fatal(err)
		}
	}

	seen := map[string]bool{}
	cursor := Cursor{}
	sizes := []int{}

	for page := 0; page < 10; page++ {
		got, err := s.History(ctx, cursor, 50)
		if err != nil {
			t.Fatal(err)
		}

		if len(got.Jobs) == 0 {
			break
		}

		sizes = append(sizes, len(got.Jobs))

		for _, job := range got.Jobs {
			if seen[job.ID] {
				t.Fatalf("row %s repeated", job.ID)
			}

			seen[job.ID] = true
		}

		if !got.HasMore {
			break
		}

		cursor = got.Next
	}

	if len(sizes) != 3 || sizes[0] != 50 || sizes[1] != 50 || sizes[2] != 20 {
		t.Fatalf("page sizes %v", sizes)
	}

	if len(seen) != 120 {
		t.Fatalf("saw %d distinct rows", len(seen))
	}

	first, _ := s.History(ctx, Cursor{}, 50)
	if first.Jobs[0].ID != "h-119" {
		t.Errorf("newest first broken: %s", first.Jobs[0].ID)
	}
}
