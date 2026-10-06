package store

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestConcurrentAccess hammers the one worker from many goroutines.
func TestConcurrentAccess(t *testing.T) {
	s := newStore(t)
	ctx := context.Background()

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			id := fmt.Sprintf("c-%02d", i)
			if err := s.SaveJob(ctx, testJob(id)); err != nil {
				t.Errorf("save %s: %v", id, err)

				return
			}

			if err := s.Checkpoint(ctx, id, 1, 2, time.Now()); err != nil {
				t.Errorf("checkpoint %s: %v", id, err)
			}

			if _, err := s.Summary(ctx); err != nil {
				t.Errorf("summary: %v", err)
			}

			if _, err := s.History(ctx, Cursor{}, 5); err != nil {
				t.Errorf("history: %v", err)
			}

			if _, err := s.ActiveJobs(ctx); err != nil {
				t.Errorf("active: %v", err)
			}
		}(i)
	}

	wg.Wait()

	counts, err := s.Summary(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if counts.Total != 20 {
		t.Errorf("saved %d rows, want 20", counts.Total)
	}
}

// TestClosedStoreRefusesWork returns ErrClosed after Close.
func TestClosedStoreRefusesWork(t *testing.T) {
	s, err := OpenMemory()
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Close(); err != nil {
		t.Fatal(err)
	}

	if err := s.Close(); err != nil {
		t.Fatal(err)
	}

	if err := s.SaveJob(context.Background(), testJob("late")); err == nil {
		t.Fatal("closed store accepted a write")
	}
}
