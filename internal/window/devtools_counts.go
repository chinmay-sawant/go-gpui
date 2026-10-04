package window

import "github.com/chinmay-sawant/gowkhtmltopdf/layout"

// devOpCounts is the per-kind operation count the panel prints.
type devOpCounts struct {
	Fill   int
	Stroke int
	Line   int
	Text   int
	Image  int
	Grid   int
}

func (c devOpCounts) total() int {
	return c.Fill + c.Stroke + c.Line + c.Text + c.Image + c.Grid
}

// devCountOps counts the operations that paint, in one pass.
func devCountOps(display *layout.Display) devOpCounts {
	var c devOpCounts
	if display == nil {
		return c
	}

	for i := range display.Ops {
		switch display.Ops[i].Kind {
		case layout.DisplayOpFillRect:
			c.Fill++
		case layout.DisplayOpStrokeRect:
			c.Stroke++
		case layout.DisplayOpLine:
			c.Line++
		case layout.DisplayOpText, layout.DisplayOpBullet:
			c.Text++
		case layout.DisplayOpImage:
			c.Image++
		case layout.DisplayOpGridRun:
			c.Grid++
		}
	}

	return c
}

// devOpsInBox counts the painted operations inside one box.
func devOpsInBox(display *layout.Display, box layout.Box) int {
	count := 0

	for i := range display.Ops {
		op := &display.Ops[i]
		if _, ok := devOpInk(op.Kind); !ok {
			continue
		}

		r := devOpRect(op, display.PixelPerPoint)
		if r.X >= box.X && r.Y >= box.Y && r.X+r.W <= box.X+box.W && r.Y+r.H <= box.Y+box.H {
			count++
		}
	}

	return count
}

// devOpRect converts an operation's bounds to page pixels. A text op carries
// its baseline in Y, so the line box starts one line height above it.
func devOpRect(op *layout.DisplayOp, ppp float64) devRect {
	x, y := op.X, op.Y
	w, h := op.W, op.H

	if op.Kind == layout.DisplayOpText || op.Kind == layout.DisplayOpBullet {
		y = op.Y + op.InkDescent - op.H
	}

	return devRect{X: x * ppp, Y: y * ppp, W: w * ppp, H: h * ppp}
}
