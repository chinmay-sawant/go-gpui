package render

import "github.com/chinmay-sawant/blinkless/layout"

// Replayable reports whether a replay painter can draw every operation in
// display the way the engine's bitmap painter would: same order, colors, and
// glyph placement. A false result means the caller should keep the bitmap
// path for this page.
//
// A CSS outline is an ordinary line or stroke op that Display.Order moves to
// the outline paint layer. Draw iterates that order, so outlines replay too.
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

	if op.Group() != nil || op.GroupBoundary() != 0 {
		return false
	}

	if op.XformSet && !op.Transform().IsIdentity() && op.Kind != layout.DisplayOpImage {
		return false
	}

	switch op.Kind {
	case layout.DisplayOpNoop, layout.DisplayOpLinkURI:
		return true
	case layout.DisplayOpFillRect:
		_, ok := FillRadii(op)

		return ok
	case layout.DisplayOpStrokeRect:
		// Top, right, bottom, and left are the only defined mask bits.
		return op.StrokeMask&^0x0F == 0
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
