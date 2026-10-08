package window

import (
	"github.com/chinmay-sawant/blinkless/layout"

	"github.com/chinmay-sawant/ownframe/internal/replay"
)

// devOpRect is one operation's painted box in canvas CSS pixels: where the
// replay draws its ink, so an outline sits on the paint. replay.PaintBounds
// owns the rule, which covers the baseline of a text run, the stroke width of
// a border, the inward geometry of a mixed border line, and the segments of a
// table grid. An op that cannot be bounded, such as a transformed run, falls
// back to its nominal box converted from points.
func devOpRect(op *layout.DisplayOp, ppp float64) devRect {
	if ppp <= 0 {
		return devRect{}
	}

	box, ok := replay.PaintBounds(op, ppp)
	if !ok || box.Empty() {
		return devNominalRect(op, ppp)
	}

	return devRect{
		X: float64(box.Min.X), Y: float64(box.Min.Y),
		W: float64(box.Dx()), H: float64(box.Dy()),
	}
}

// devNominalRect converts an op's raw box from points to CSS pixels. A text
// op carries its baseline in Y, so its nominal box starts one line height
// above it.
func devNominalRect(op *layout.DisplayOp, ppp float64) devRect {
	x, y, w, h := op.X, op.Y, op.W, op.H

	if op.Kind == layout.DisplayOpText || op.Kind == layout.DisplayOpBullet {
		y = op.Y + op.InkDescent - op.H
	}

	return devRect{X: x / ppp, Y: y / ppp, W: w / ppp, H: h / ppp}
}
