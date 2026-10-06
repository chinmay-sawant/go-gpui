package wire

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// TestRestartKeepsData opens the same data directory twice: the seed does
// not duplicate, history keeps its rows, and the theme survives.
func TestRestartKeepsData(t *testing.T) {
	ctx := context.Background()
	dir := t.TempDir()

	first := newBackend(t, Config{DataDir: dir, Dummy: true})
	app, err := ui.New(ctx, first)
	if err != nil {
		first.Close()
		t.Fatalf("ui.New: %v", err)
	}

	tick(t, app, 30)

	before := app.View().Summary.Completed
	if before < 1 {
		t.Fatalf("first session summary = %+v", app.View().Summary)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	clickVisible(t, app, "theme")

	if err := app.Close(); err != nil {
		t.Fatalf("first Close: %v", err)
	}

	second := newBackend(t, Config{DataDir: dir, Dummy: true})
	app2, err := ui.New(ctx, second)
	if err != nil {
		second.Close()
		t.Fatalf("second ui.New: %v", err)
	}

	defer func() {
		if err := app2.Close(); err != nil {
			t.Errorf("second Close: %v", err)
		}
	}()

	if dark, err := second.Dark(); err != nil || !dark {
		t.Fatalf("theme after restart: dark=%v err=%v", dark, err)
	}

	tick(t, app2, 30)

	after := app2.View().Summary.Completed
	if after != before {
		t.Fatalf("completed count changed across restart: %d then %d", before, after)
	}
}
