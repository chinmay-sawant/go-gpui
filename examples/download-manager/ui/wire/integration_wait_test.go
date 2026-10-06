package wire

import (
	"os"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

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

// waitState ticks until the job with url reaches want, and returns its ID.
func waitState(t *testing.T, app *ui.App, url string, want ui.State, timeout time.Duration) string {
	t.Helper()
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		tick(t, app, 3)

		for _, row := range app.View().Active {
			if row.URL == url && row.State == want {
				return row.ID
			}
		}

		for _, row := range app.View().History {
			if row.URL == url && row.State == want {
				return row.ID
			}
		}
	}

	for _, row := range app.View().Active {
		t.Logf("active row: %+v", row)
	}

	t.Fatalf("the job never reached %v; notice=%q", want, app.View().Notice)

	return ""
}
