package transfer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/fixture"
)

// TestHTTPStall fails a slow trickle under a short window.
func TestHTTPStall(t *testing.T) {
	srv := fixture.NewServer()
	srv.SlowPause = 300 * time.Millisecond
	defer srv.Close()

	tr := NewHTTP(HTTPOptions{StallTimeout: 40 * time.Millisecond})
	dest, partial := httpPaths(t)

	_, err := tr.Download(context.Background(), Request{
		URL: srv.URL + fixture.PathSlow, Dest: dest, Partial: partial,
	}, nil)
	if !errors.Is(err, ErrStall) {
		t.Fatalf("want ErrStall, got %v", err)
	}
}
