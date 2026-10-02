package window

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const badgeLabel = "bitmap fallback"

var (
	badgeFace = text.NewGoXFace(basicfont.Face7x13)
	badgeInk  = color.RGBA{R: 0x1c, G: 0x19, B: 0x15, A: 0xe6}
)

// badgeRect returns the top-right box that holds label on a screenW-wide frame.
// The box keeps an 8 px margin from the right and top edges and 4 px of padding.
func badgeRect(screenW float64, label string) (x, y, width, height float64) {
	labelW, labelH := text.Measure(label, badgeFace, 0)

	width = labelW + 8
	height = labelH + 8
	x = screenW - width - 8
	y = 8

	return x, y, width, height
}

// drawBadge paints the bitmap fallback notice on top of the frame.
func (s *shell) drawBadge(screen *ebiten.Image) {
	x, y, width, height := badgeRect(float64(s.screenW), badgeLabel)

	vector.FillRect(
		screen,
		float32(x), float32(y), float32(width), float32(height),
		badgeInk, false,
	)

	var op text.DrawOptions
	op.GeoM.Translate(x+4, y+4)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, badgeLabel, badgeFace, &op)
}
