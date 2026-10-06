package scheduler

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestQueueFull rejects a job when the waiting list is full.
func TestQueueFull(t *testing.T) {
	gate := newGate()
	eng, _ := newTestEngine(t, gate, Options{Workers: 1, Queue: 1})
	eng.Start(context.Background())

	first := add(t, eng, t.TempDir(), "https://example.invalid/first.bin")
	waitStarted(t, gate, first.ID)

	add(t, eng, t.TempDir(), "https://example.invalid/second.bin")

	_, err := eng.Add(context.Background(), AddRequest{
		URL: "https://example.invalid/third.bin", Dir: t.TempDir(),
	})
	if !errors.Is(err, ErrQueueFull) {
		t.Fatalf("want ErrQueueFull, got %v", err)
	}
}

// TestShutdownCancelsAndPersists stops a live transfer within budget.
func TestShutdownCancelsAndPersists(t *testing.T) {
	gate := newGate()
	eng, st := newTestEngine(t, gate, Options{
		Workers: 1, ShutdownBudget: 500 * time.Millisecond,
	})
	eng.Start(context.Background())

	job := add(t, eng, t.TempDir(), "https://example.invalid/close.bin")
	waitStarted(t, gate, job.ID)

	start := time.Now()

	if err := eng.Close(context.Background()); err != nil {
		t.Fatal(err)
	}

	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Errorf("close took %v", elapsed)
	}

	row, err := st.Job(context.Background(), job.ID)
	if err != nil {
		t.Fatal(err)
	}

	if row.State != domain.StatePaused {
		t.Errorf("shutdown state %s, want paused", row.State)
	}

	if err := eng.Close(context.Background()); err != nil {
		t.Errorf("second close: %v", err)
	}

	if _, err := eng.Add(context.Background(), AddRequest{
		URL: "https://example.invalid/late.bin", Dir: t.TempDir(),
	}); !errors.Is(err, ErrClosed) {
		t.Errorf("add after close: %v", err)
	}
}
