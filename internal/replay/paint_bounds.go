package replay

import (
	"image"

	"github.com/chinmay-sawant/blinkless/layout"
)

// PaintBounds returns the op's painted box in CSS pixels, rounded outward,
// with no slack. opBounds is the same box grown by one pixel per side for the
// replay's dirty-rect filter; a reader that wants the ink itself, such as the
// window's inspector, uses this.
//
// The box follows the paint: a text or bullet run is bounded by its face
// ascent and ink descent around the baseline in Y, a line by its stroke width
// and inward geometry, a stroke rect by half its width outside the box, and a
// grid run by the union of its segments.
//
// ok is false when the box cannot be bounded, such as a transformed or
// rotated run; the caller must treat such an op as unbounded.
func PaintBounds(op *layout.DisplayOp, ppt float64) (image.Rectangle, bool) {
	box, ok := opBounds(op, ppt)
	if !ok || box.Empty() {
		return box, ok
	}

	// opBounds adds boundsPad after rounding outward, so taking the same
	// whole pixels back is exact.
	return box.Inset(boundsPad), ok
}
