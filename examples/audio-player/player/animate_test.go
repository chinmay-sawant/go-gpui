package player

import (
	"context"
	"math"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/frame"
)

func TestTickAnimatesFrame(t *testing.T) {
	app, voice := newAudioApp(t)
	laid := eqHeights(t, app)

	waitLoaded(t, app)
	voice.pos = voice.dur / 2

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	d := app.Page().Display()
	seek, ok := boxByID(app.Boxes(), "seek")
	if !ok {
		t.Fatal("no seek box")
	}

	fill := frame.Fill(d, seek, accent)
	if fill == nil {
		t.Fatal("no seek fill")
	}

	_, _, w, _ := frame.BoxUnits(d, seek)
	if fill.W <= 0 || math.Abs(fill.W-w/2) > 1 {
		t.Fatalf("seek fill W = %v, want %v", fill.W, w/2)
	}

	if sameHeights(laid, eqHeights(t, app)) {
		t.Fatal("no eq bar moved")
	}

	elapsed, ok := boxByID(app.Boxes(), "elapsed")
	if !ok {
		t.Fatal("no elapsed box")
	}

	want := clock(int(voice.pos.Seconds()))
	if op := frame.Text(d, elapsed); op == nil || op.Text != want {
		t.Fatalf("elapsed text = %+v, want %q", op, want)
	}
}

func TestPausedEqStaysStill(t *testing.T) {
	app, _ := newAudioApp(t)
	waitLoaded(t, app)

	clickBox(t, app, "play") // pause

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	first := eqHeights(t, app)

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if !sameHeights(first, eqHeights(t, app)) {
		t.Fatal("eq bars moved while paused")
	}
}
