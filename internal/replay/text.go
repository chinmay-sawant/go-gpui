package replay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/chinmay-sawant/blinkless/layout"
)

// drawText draws one shaped run at its baseline. The engine stores X and Y as
// the baseline origin, while Ebiten's text origin is the top of the line, so
// the draw moves up by the face ascent. Fake bold is a one-pixel double
// strike, as in the engine's own painter. A letter-spaced run is drawn glyph
// by glyph with the extra advance.
func drawText(dst *ebiten.Image, op *layout.DisplayOp, dx, dy float64) {
	face := textFace(op)
	if face == nil || op.Text == "" {
		return
	}

	content := layout.DisplayTransformText(op.Text, op.TextTransformValue())
	x := op.X*pxPerPt + dx
	y := op.Y*pxPerPt + dy - face.Metrics().HAscent

	if op.LetterSpacing == 0 {
		drawRun(dst, content, face, x, y, op)

		return
	}

	for _, r := range content {
		glyph := string(r)
		drawRun(dst, glyph, face, x, y, op)
		width, _ := text.Measure(glyph, face, 0)
		x += width + op.LetterSpacing*pxPerPt
	}
}

func drawRun(dst *ebiten.Image, content string, face *text.GoTextFace, x, y float64, op *layout.DisplayOp) {
	options := &text.DrawOptions{}
	options.GeoM.Translate(x, y)
	options.ColorScale.ScaleWithColor(rgba(op))

	text.Draw(dst, content, face, options)

	if layout.DisplayFakeBold(op) {
		options.GeoM.Translate(1, 0)
		text.Draw(dst, content, face, options)
	}
}
