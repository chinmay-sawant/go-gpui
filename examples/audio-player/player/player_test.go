package player

import (
	"bytes"
	"context"
	"image/color"
	"image/png"
	"testing"
)

func TestNewDrawsReplayablePlayer(t *testing.T) {
	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	if app.Page().Display() == nil {
		t.Fatal("the player fell back to the bitmap path")
	}

	if len(app.PNG()) == 0 {
		t.Fatal("no PNG after Redraw")
	}
}

// TestSampleViewRendersPlaceholders pins the offline look: the sample covers
// are embedded SVGs, so the frame must carry their fill colors.
func TestSampleViewRendersPlaceholders(t *testing.T) {
	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	assertReplayable(t, app)

	img, err := png.Decode(bytes.NewReader(app.PNG()))
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	placeholder := color.RGBA{R: 221, G: 214, B: 254, A: 255}
	if !hasColor(img, placeholder, 6) {
		t.Fatal("no cover-1 placeholder pixels in the sample frame")
	}
}
