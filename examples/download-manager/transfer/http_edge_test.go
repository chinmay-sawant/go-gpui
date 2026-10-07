package transfer

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/fixture"
)

// TestHTTPChecksumMismatch removes the corrupt partial.
func TestHTTPChecksumMismatch(t *testing.T) {
	srv := fixture.NewServer()
	defer srv.Close()

	dest, partial := httpPaths(t)

	_, err := newHTTP(t).Download(context.Background(), Request{
		URL: srv.URL + fixture.PathOK, Dest: dest, Partial: partial,
		Checksum: "sha256:00",
	}, nil)
	if !errors.Is(err, ErrChecksum) {
		t.Fatalf("want ErrChecksum, got %v", err)
	}

	if _, statErr := os.Stat(partial); !os.IsNotExist(statErr) {
		t.Error("corrupt partial kept")
	}
}

// TestHTTPRejectsContentEncoding refuses an encoded body.
func TestHTTPRejectsContentEncoding(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Length", "4")
		_, _ = w.Write([]byte("abcd"))
	}))
	defer srv.Close()

	dest, partial := httpPaths(t)

	_, err := newHTTP(t).Download(context.Background(), Request{
		URL: srv.URL, Dest: dest, Partial: partial,
	}, nil)
	if !errors.Is(err, ErrUnsupportedResume) {
		t.Fatalf("want ErrUnsupportedResume, got %v", err)
	}
}
