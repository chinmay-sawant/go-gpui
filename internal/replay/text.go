package replay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// drawText draws one shaped run at its baseline. The engine stores X and Y as
// the baseline origin, while Ebiten's text origin is the top of the line, so
// the draw moves up by the face ascent. Fake bold is a one-pixel double
// strike, as in the engine's own painter.
func drawText(dst *ebiten.Image, op *layout.DisplayOp, dx, dy float64) {
	face := textFace(op)
	if face == nil || op.Text == "" {
		return
	}

	content := layout.DisplayTransformText(op.Text, op.TextTransformValue())
	options := &text.DrawOptions{}
	options.GeoM.Translate(op.X*pxPerPt+dx, op.Y*pxPerPt+dy-face.Metrics().HAscent)
	options.ColorScale.ScaleWithColor(rgba(op))

	text.Draw(dst, content, face, options)

	if layout.DisplayFakeBold(op) {
		options.GeoM.Translate(1, 0)
		text.Draw(dst, content, face, options)
	}
}
