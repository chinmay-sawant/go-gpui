package music

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestOpenverseResolveUsesCache(t *testing.T) {
	t.Parallel()

	var downloads atomic.Int32

	srv := ovServer(t, &downloads)
	defer srv.Close()

	o := &Openverse{Base: srv.URL, CacheDir: t.TempDir()}

	if _, err := o.Resolve(context.Background(), "jazz", 0); err != nil {
		t.Fatalf("first Resolve: %v", err)
	}

	if _, err := o.Resolve(context.Background(), "jazz", 0); err != nil {
		t.Fatalf("second Resolve: %v", err)
	}

	if got := downloads.Load(); got != 1 {
		t.Fatalf("downloads = %d, want 1", got)
	}
}

func TestOpenverseSearchFailure(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer srv.Close()

	o := &Openverse{Base: srv.URL}

	if _, err := o.Resolve(context.Background(), "jazz", 0); err == nil {
		t.Fatal("Resolve accepted a failed search")
	}
}

func TestLibraryFallsBackToDemoTune(t *testing.T) {
	t.Parallel()

	lib := &Library{Primary: &Openverse{Base: "http://127.0.0.1:1"}}
	clip, err := lib.Resolve(context.Background(), "jazz", 0)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	if clip.License != "generated" {
		t.Fatalf("clip = %+v", clip)
	}
}
