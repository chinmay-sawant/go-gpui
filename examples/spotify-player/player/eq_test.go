package player

import (
	"context"
	"testing"
)

func TestPausedEqStaysStatic(t *testing.T) {
	app, _ := newFakeApp(t)
	redraw(t, app)

	click(t, app, "play")
	waitAudio(t, app)
	click(t, app, "play")

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	first := eqHeights(app)

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	second := eqHeights(app)

	if len(first) == 0 || len(first) != len(second) {
		t.Fatalf("heights = %v / %v", first, second)
	}

	for i := range first {
		if first[i] != second[i] {
			t.Fatalf("bar %d moved while paused: %v -> %v", i, first[i], second[i])
		}
	}
}
