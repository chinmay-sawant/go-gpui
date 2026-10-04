package window

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// drawDevPanel paints the panel background and one line per row.
func (s *shell) drawDevPanel(screen *ebiten.Image, lines []string) {
	r := s.dev.panel
	if r.W <= 0 || r.H <= 0 {
		return
	}

	vector.FillRect(screen, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), badgeInk, false)

	lineH := devLineHeight()

	for i, line := range lines {
		var op text.DrawOptions
		op.GeoM.Translate(r.X+devPanelPad, r.Y+devPanelPad+float64(i)*lineH)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, line, badgeFace, &op)
	}
}

// devLineHeight is the badge face line height plus a little leading.
func devLineHeight() float64 {
	_, h := text.Measure("Ag", badgeFace, 0)

	return h + devLineGap
}
