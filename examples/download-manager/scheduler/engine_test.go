package scheduler

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestAddCompletes persists the job and finishes it.
func TestAddCompletes(t *testing.T) {
	eng, st := newTestEngine(t, instantTransport{}, Options{})
	eng.Start(context.Background())

	job := add(t, eng, t.TempDir(), "https://example.invalid/a.bin")
	done := waitState(t, eng, job.ID, domain.StateCompleted)

	if done.Done != 7 || done.Total != 7 {
		t.Errorf("outcome %+v", done)
	}

	row, err := st.Job(context.Background(), job.ID)
	if err != nil {
		t.Fatal(err)
	}

	if row.State != domain.StateCompleted || row.Done != 7 {
		t.Errorf("persisted row %+v", row)
	}

	if got, ok := eng.Job(job.ID); !ok || got.State != domain.StateCompleted {
		t.Errorf("live job %+v ok=%v", got, ok)
	}
}

// TestUnknownJobControl rejects IDs the engine never saw.
func TestUnknownJobControl(t *testing.T) {
	eng, _ := newTestEngine(t, instantTransport{}, Options{})
	ctx := context.Background()

	for _, err := range []error{
		eng.Pause(ctx, "nope"),
		eng.Resume(ctx, "nope"),
		eng.Cancel(ctx, "nope"),
		eng.Retry(ctx, "nope"),
	} {
		if err == nil {
			t.Error("unknown job accepted")
		}
	}

	if _, ok := eng.Job("nope"); ok {
		t.Error("unknown job found")
	}
}
