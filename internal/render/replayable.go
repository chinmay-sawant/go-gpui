package render

import "github.com/chinmay-sawant/gowkhtmltopdf/layout"

// Replayable reports whether a replay painter can draw every operation in
// display the way the engine's bitmap painter would: same order, colors, and
// glyph placement. A false result means the caller should keep the bitmap
// path for this page.
func Replayable(display *layout.Display) bool {
	if display == nil {
		return false
	}

	for i := range display.Ops {
		if !replayableOp(&display.Ops[i]) {
			return false
		}
	}

	return true
}

func replayableOp(op *layout.DisplayOp) bool {
	if op.BlendModeName() != "" && op.BlendModeName() != "normal" {
		return false
	}

	if op.Group() != nil || op.GroupBoundary() != 0 || op.Outline() {
		return false
	}

	if op.XformSet && !op.Transform().IsIdentity() {
		return false
	}

	switch op.Kind {
	case layout.DisplayOpNoop, layout.DisplayOpLinkURI:
		return true
	case layout.DisplayOpFillRect:
		_, ok := FillRadii(op)

		return ok
	case layout.DisplayOpStrokeRect:
		if op.StrokeMask != 0 {
			return false
		}

		_, ok := FillRadii(op)

		return ok
	case layout.DisplayOpImage:
		data, _, _ := op.ImageBytes()

		return data != nil
	case layout.DisplayOpLine:
		return true
	case layout.DisplayOpText, layout.DisplayOpBullet:
		return replayableText(op)
	case layout.DisplayOpGridRun:
		return op.Grid != nil
	default:
		return false
	}
}

func replayableText(op *layout.DisplayOp) bool {
	if op.Font == nil || op.RotateDeg != 0 || op.FakeOblique {
		return false
	}

	return op.FontFeatures() == "" && op.TextAutospaceGap() == 0
}
