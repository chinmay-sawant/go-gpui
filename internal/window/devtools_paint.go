package window

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

var (
	devHoverInk = color.RGBA{R: 0xe5, G: 0x48, B: 0x4d, A: 0xff}
	devPinInk   = color.RGBA{R: 0x1a, G: 0x56, B: 0xdb, A: 0xff}
)

// drawDevPick outlines the hovered and pinned boxes, one colour each, with a
// label over each.
func (s *shell) drawDevPick(screen *ebiten.Image) {
	if s.dev.haveHov {
		s.devOutline(screen, s.dev.hovered, devHoverInk)
	}

	if s.dev.havePin {
		s.devOutline(screen, s.dev.pinned, devPinInk)
	}
}

// devOutline strokes one box and labels it.
func (s *shell) devOutline(screen *ebiten.Image, box layout.Box, ink color.RGBA) {
	r := s.devScreen(devRect{X: box.X, Y: box.Y, W: box.W, H: box.H})
	vector.StrokeRect(screen, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), 2, ink, false)
	s.devLabel(screen, r, devBoxLabel(box))
}

// devBoxLabel is the tag, id, and size line drawn over a box.
func devBoxLabel(box layout.Box) string {
	name := box.Tag
	if box.ID != "" {
		name += "#" + box.ID
	}

	return fmt.Sprintf("%s  %gx%g", name, box.W, box.H)
}

// devLabel draws a box label above the outline, or below it when the top
// edge would clip.
func (s *shell) devLabel(screen *ebiten.Image, r devRect, label string) {
	w, h := text.Measure(label, badgeFace, 0)
	x, y := r.X, r.Y-h-4

	if y < 0 {
		y = r.Y + r.H + 4
	}

	vector.FillRect(screen, float32(x), float32(y), float32(w+8), float32(h+4), badgeInk, false)

	var op text.DrawOptions
	op.GeoM.Translate(x+4, y+2)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, label, badgeFace, &op)
}
