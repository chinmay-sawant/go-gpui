package replay

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/blinkless/layout"
)

// drawGrid replays one collapsed table grid as its line segments, in
// emission order. A segment carries its own color and width; the engine's
// grid painter draws every segment opaque, ignoring the owner's alpha.
func drawGrid(dst *ebiten.Image, op *layout.DisplayOp, dx, dy float64) {
	if op.Grid == nil {
		return
	}

	for i := range op.Grid.Segs {
		seg := &op.Grid.Segs[i]

		line := layout.DisplayOp{
			ID:   op.ID,
			Kind: layout.DisplayOpLine,
			X:    seg.X, Y: seg.Y, W: seg.W, H: seg.H,
			Width: seg.Width, R: seg.R, G: seg.G, B: seg.B,
			LineInset: seg.LineInset,
		}

		drawLine(dst, &line, dx, dy)
	}
}
