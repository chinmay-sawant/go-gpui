package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// acceptDisplay lays source out on a small frame, so each assertion is about
// a real page rather than a hand-built op list.
func acceptDisplay(t *testing.T, source string) *render.Display {
	t.Helper()

	display, err := render.DisplayList(context.Background(), source, 320, 240)
	if err != nil {
		t.Fatal(err)
	}

	return display
}

func acceptKinds(display *render.Display) map[render.Kind]int {
	got := map[render.Kind]int{}

	for _, op := range display.Ops {
		got[op.Kind]++
	}

	return got
}

// TestReplayAcceptsGradientCard is a page that previously fell back: a
// rounded border over a linear gradient. The engine rasterizes the gradient
// to image ops and strokes the border, so the gate must accept both.
func TestReplayAcceptsGradientCard(t *testing.T) {
	t.Parallel()

	display := acceptDisplay(t, `<div style="width:240px;padding:16px;`+
		`border:2px solid #333;border-radius:12px;`+
		`background-image:linear-gradient(to right,#e8f0ff,#ffffff)">`+
		`<h1>Sign in</h1><p>Use your account to continue.</p></div>`)

	if !render.Replayable(display) {
		t.Fatal("gradient card should replay")
	}

	kinds := acceptKinds(display)
	if kinds[render.OpStrokeRect] == 0 {
		t.Error("no stroke rect for the rounded border")
	}

	if kinds[render.OpImage] == 0 {
		t.Error("no image op for the linear gradient")
	}
}

// TestReplayAcceptsPlainFillsAndText covers the base case: solid fills and
// shaped text replay without a raster fallback.
func TestReplayAcceptsPlainFillsAndText(t *testing.T) {
	t.Parallel()

	display := acceptDisplay(t,
		`<div style="width:120px;height:40px;background:#eee"></div>`+
			`<p>plain body copy</p>`)

	if !render.Replayable(display) {
		t.Fatal("plain fills and text should replay")
	}

	kinds := acceptKinds(display)
	if kinds[render.OpFillRect] == 0 || kinds[render.OpText] == 0 {
		t.Fatalf("kinds = %v, want fill and text", kinds)
	}
}
