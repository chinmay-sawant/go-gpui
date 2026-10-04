package window

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// TestDevDrawOverlaySmoke draws the overlay into an offscreen image in
// normal and stretched modes. The outline coordinates come from the same
// conversion the draw path uses.
func TestDevDrawOverlaySmoke(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.display = devTestDisplay()
	s.dev.ops = true
	s.dev.haveHov = true
	s.dev.hovered = d.boxes[0]
	s.dev.havePin = true
	s.dev.pinned = d.boxes[0]

	img := ebiten.NewImage(1920, 1080)
	s.drawDevTools(img)

	got := s.devScreen(devRect{X: 10, Y: 10, W: 100, H: 40})
	want := devRect{X: 24, Y: 18, W: 240, H: 72}

	if got != want {
		t.Fatalf("stretched outline = %+v, want %+v", got, want)
	}

	s.screenW, s.screenH = 800, 600
	s.drawDevTools(img)

	got = s.devScreen(devRect{X: 10, Y: 10, W: 100, H: 40})
	want = devRect{X: 10, Y: 10, W: 100, H: 40}

	if got != want {
		t.Fatalf("normal outline = %+v, want %+v", got, want)
	}

	// A page with no boxes still draws its panel.
	d.boxes = nil
	s.dev.haveHov, s.dev.havePin = false, false
	s.drawDevTools(img)

	if s.dev.panel.W <= 0 || s.dev.panel.H <= 0 {
		t.Fatalf("panel = %+v, want a drawn panel", s.dev.panel)
	}
}

// devTestDisplay is a small display list with one fill and one text run.
func devTestDisplay() *layout.Display {
	return &layout.Display{
		Ops: []layout.DisplayOp{
			{Kind: layout.DisplayOpFillRect, X: 0, Y: 0, W: 100, H: 50},
			{Kind: layout.DisplayOpText, X: 10, Y: 30, W: 80, H: 16, Text: "hi"},
			{Kind: layout.DisplayOpNoop, X: 0, Y: 0, W: 1, H: 1},
		},
		Order:         []int{0, 1, 2},
		Width:         800,
		Height:        600,
		PixelPerPoint: 0.75,
	}
}
