package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestProgressCheckpointed writes bytes before the transfer ends.
func TestProgressCheckpointed(t *testing.T) {
	gate := newGate()
	eng, st := newTestEngine(t, gate, Options{
		Workers: 1, ProgressEvery: 10 * time.Millisecond,
	})
	eng.Start(context.Background())

	job := add(t, eng, t.TempDir(), "https://example.invalid/progress.bin")
	waitStarted(t, gate, job.ID)

	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) {
		row, err := st.Job(context.Background(), job.ID)
		if err != nil {
			t.Fatal(err)
		}

		if row.Done == 5 {
			if row.State != domain.StateRunning {
				t.Errorf("state %s while running", row.State)
			}

			return
		}

		time.Sleep(5 * time.Millisecond)
	}

	t.Fatal("progress never checkpointed")
}
