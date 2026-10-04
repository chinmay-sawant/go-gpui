package window

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// TestDevDrawOverlaySmoke draws the overlay into an offscreen image in
// normal and stretched modes. The stroke hook records the computed screen
// rects, so the test finds each outline without reading pixels back.
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

	var strokes []devRect
	s.dev.stroke = func(_ *ebiten.Image, r devRect, _ color.RGBA) {
		strokes = append(strokes, r)
	}

	img := ebiten.NewImage(1920, 1080)
	s.drawDevTools(img)

	box := d.boxes[0]
	want := s.devScreen(devRect{X: box.X, Y: box.Y, W: box.W, H: box.H})
	if len(strokes) != 2 || strokes[0] != want || strokes[1] != want {
		t.Fatalf("stretched outlines = %v, want %v twice", strokes, want)
	}

	s.screenW, s.screenH = 800, 600
	s.drawDevTools(img)

	want = s.devScreen(devRect{X: box.X, Y: box.Y, W: box.W, H: box.H})
	if len(strokes) != 4 || strokes[2] != want || strokes[3] != want {
		t.Fatalf("normal outlines = %v, want %v twice", strokes[2:], want)
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
