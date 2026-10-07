package ui

import (
	"context"
	"testing"
)

// running builds one running row for tests.
func running(id string, done, total int64) Row {
	return Row{ID: id, Name: id, State: StateRunning, Done: done, Total: total, Speed: 1 << 20}
}

func TestTickAppliesActiveRows(t *testing.T) {
	app, back := newTestApp(t)
	ctx := context.Background()

	back.send(Update{Kind: UpdateActive, Active: []Row{running("j1", 10, 100), running("j2", 20, 100)}})

	if err := app.Tick(ctx); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if len(app.view.Active) != 2 {
		t.Fatalf("active = %d, want 2", len(app.view.Active))
	}

	if _, ok := findBox(app.Boxes(), "track-1"); !ok {
		t.Fatal("no track-1 box after the active update")
	}

	if _, ok := findBox(app.Boxes(), "fill-1"); !ok {
		t.Fatal("no fill-1 box after the active update")
	}
}

func TestTickProgressIsPaintOnly(t *testing.T) {
	app, back := newTestApp(t)
	ctx := context.Background()

	back.send(Update{Kind: UpdateActive, Active: []Row{running("j1", 10, 100)}})

	if err := app.Tick(ctx); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	redraws := app.view.Stats.Redraws
	row := running("j1", 60, 100)
	back.send(Update{Kind: UpdateProgress, Row: &row})

	if err := app.Tick(ctx); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if app.view.Stats.Redraws != redraws {
		t.Fatal("a progress update forced a redraw")
	}

	if app.view.Stats.Paints == 0 {
		t.Fatal("no paint-only tick counted")
	}

	if len(app.bindings.fill) != 1 || app.bindings.fill[0] == nil {
		t.Fatal("fill operation not bound")
	}

	if app.bindings.fill[0].W <= 0 {
		t.Fatalf("fill width = %v", app.bindings.fill[0].W)
	}
}
