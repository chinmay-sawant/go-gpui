package wire

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// newBackend opens a backend or fails the test.
func newBackend(t *testing.T, cfg Config) *Backend {
	t.Helper()

	backend, err := New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("wire.New: %v", err)
	}

	return backend
}

// waitHistory ticks until the job with url reaches want on the history
// page.
func waitHistory(t *testing.T, app *ui.App, url string, want ui.State, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		tick(t, app, 3)

		for _, row := range app.View().History {
			if row.URL == url && row.State == want {
				return
			}
		}
	}

	for _, row := range app.View().Active {
		t.Logf("active row: %+v", row)
	}

	t.Fatalf("the job never reached %v in history; notice=%q", want, app.View().Notice)
}

// tick runs n UI ticks with a short pause.
func tick(t *testing.T, app *ui.App, n int) {
	t.Helper()
	ctx := context.Background()

	for i := 0; i < n; i++ {
		if err := app.Tick(ctx); err != nil {
			t.Fatalf("Tick: %v", err)
		}

		time.Sleep(10 * time.Millisecond)
	}
}

// clickVisible clicks the action box in the current drawing.
func clickVisible(t *testing.T, app *ui.App, action string) {
	t.Helper()
	ctx := context.Background()

	for _, box := range app.Boxes() {
		if box.Action != action {
			continue
		}

		if err := app.Click(ctx, box.X+box.W/2, box.Y+box.H/2); err != nil {
			t.Fatalf("Click %s: %v", action, err)
		}

		return
	}

	t.Fatalf("no box with action %q", action)
}

// waitFile waits for a file the worker writes.
func waitFile(t *testing.T, path string, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		if _, err := os.Stat(path); err == nil {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf("file %s was not written", path)
}
