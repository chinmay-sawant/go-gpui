package scheduler

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestRetryAfterFailure queues a failed job again.
func TestRetryAfterFailure(t *testing.T) {
	gate := newGate()
	eng, _ := newTestEngine(t, gate, Options{Workers: 1})
	eng.Start(context.Background())

	gate.setError(context.DeadlineExceeded)
	job := add(t, eng, t.TempDir(), "https://example.invalid/retry.bin")
	waitStarted(t, gate, job.ID)
	gate.open()

	failed := waitState(t, eng, job.ID, domain.StateFailed)
	if failed.Error == "" {
		t.Error("failure carried no error text")
	}

	if failed.Attempts != 1 {
		t.Errorf("attempts %d", failed.Attempts)
	}

	gate.setError(nil)

	if err := eng.Retry(context.Background(), job.ID); err != nil {
		t.Fatal(err)
	}

	done := waitState(t, eng, job.ID, domain.StateCompleted)
	if done.Attempts != 2 || done.Error != "" {
		t.Errorf("retried job %+v", done)
	}
}
