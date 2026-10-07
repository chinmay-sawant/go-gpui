package transfer

import (
	"context"
	"os"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/fixture"
)

// TestHTTPChunkedUnknownLength hides Content-Length.
func TestHTTPChunkedUnknownLength(t *testing.T) {
	srv := fixture.NewServer()
	defer srv.Close()

	body, out := fetchFull(t, newHTTP(t), srv.URL+fixture.PathChunked)

	if out.Total != Unknown {
		t.Errorf("total %d, want Unknown", out.Total)
	}

	if len(body) != 64<<10 {
		t.Errorf("body %d bytes", len(body))
	}
}

// TestHTTPInterrupt keeps the partial after a dropped connection.
func TestHTTPInterrupt(t *testing.T) {
	srv := fixture.NewServer()
	defer srv.Close()

	dest, partial := httpPaths(t)

	_, err := newHTTP(t).Download(context.Background(), Request{
		URL: srv.URL + fixture.PathInterrupt, Dest: dest, Partial: partial,
	}, nil)
	if err == nil {
		t.Fatal("interrupted transfer succeeded")
	}

	info, statErr := os.Stat(partial)
	if statErr != nil {
		t.Fatalf("partial not kept: %v", statErr)
	}

	if info.Size() != 32<<10 {
		t.Errorf("partial %d bytes", info.Size())
	}

	if _, statErr := os.Stat(dest); !os.IsNotExist(statErr) {
		t.Error("interrupted transfer finalized")
	}
}
