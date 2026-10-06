package ui

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/storage"
)

func TestStoreSeedsOnceAndPersistsTheme(t *testing.T) {
	ctx := context.Background()
	dir := t.TempDir()

	st, err := storage.Open(dir)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = st.Close() })

	fx := collector.DummyFixture(1, time.Unix(1700000000, 0),
		collector.DefaultHistory, collector.DefaultHistoryStep)

	seeded, err := st.Seed(ctx, fx)
	if err != nil || !seeded {
		t.Fatalf("first seed = %v, %v", seeded, err)
	}

	seeded, err = st.Seed(ctx, fx)
	if err != nil || seeded {
		t.Fatalf("second seed = %v, %v (must be idempotent)", seeded, err)
	}

	store := NewCollectorStore(st)

	if settings, err := store.LoadSettings(ctx); err != nil || settings.Dark {
		t.Fatalf("fresh settings = %+v, %v", settings, err)
	}

	if err := store.SaveSettings(ctx, Settings{Dark: true}); err != nil {
		t.Fatal(err)
	}

	if settings, err := store.LoadSettings(ctx); err != nil || !settings.Dark {
		t.Fatalf("saved settings = %+v, %v", settings, err)
	}

	// Reopen the same directory: the fixture version and theme survive.
	if err := st.Close(); err != nil {
		t.Fatal(err)
	}

	reopened, err := storage.Open(dir)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = reopened.Close() })

	if v, err := reopened.FixtureVersion(ctx); err != nil || v != collector.SeedVersion {
		t.Fatalf("fixture version after reopen = %d, %v", v, err)
	}

	if settings, err := NewCollectorStore(reopened).LoadSettings(ctx); err != nil || !settings.Dark {
		t.Fatalf("theme after reopen = %+v, %v", settings, err)
	}
}
