package music

import (
	"context"
	"sync/atomic"
	"testing"
)

func TestOpenverseResolvePicksAndFilters(t *testing.T) {
	t.Parallel()

	var downloads atomic.Int32

	srv := ovServer(t, &downloads)
	defer srv.Close()

	o := &Openverse{Base: srv.URL, CacheDir: t.TempDir()}

	first, err := o.Resolve(context.Background(), "jazz", 0)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	if first.Title != "First" || first.Creator != "Ann" || first.License != "CC BY-SA 3.0" {
		t.Fatalf("clip = %+v", first)
	}

	second, err := o.Resolve(context.Background(), "jazz", 1)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	if second.Title != "Second" || second.License != "CC0 1.0" {
		t.Fatalf("clip = %+v", second)
	}

	if got := downloads.Load(); got != 2 {
		t.Fatalf("downloads = %d, want 2", got)
	}
}
