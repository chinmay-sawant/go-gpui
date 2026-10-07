package transfer

import (
	"context"
	"path/filepath"
	"testing"
	"time"
)

// newHTTP builds a transport with test-sized timeouts.
func newHTTP(t *testing.T) *HTTP {
	t.Helper()

	return NewHTTP(HTTPOptions{
		ConnectTimeout:  time.Second,
		ResponseTimeout: time.Second,
		StallTimeout:    2 * time.Second,
	})
}

// httpPaths returns a fresh dest and partial pair for HTTP tests.
func httpPaths(t *testing.T) (string, string) {
	t.Helper()
	dir := t.TempDir()
	dest := filepath.Join(dir, "http.bin")

	return dest, PartialPath(dest)
}

// fetchFull downloads url and returns the body and outcome.
func fetchFull(t *testing.T, tr *HTTP, url string) ([]byte, Outcome) {
	t.Helper()
	dest, partial := httpPaths(t)

	out, err := tr.Download(context.Background(), Request{
		URL: url, Dest: dest, Partial: partial,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	return readAll(t, dest), out
}
