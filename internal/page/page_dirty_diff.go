package page

import (
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// maxDirtyOps caps the changed operations before the whole frame repaints.
// Merging many small rects costs more than a full replay.
const maxDirtyOps = 8

// diffDisplay returns the region covering the operations that differ between
// prev and next. full reports a difference too scattered to merge.
func diffDisplay(prev, next *layout.Display) (image.Rectangle, bool) {
	if prev.Width != next.Width || prev.Height != next.Height {
		return image.Rectangle{}, true
	}

	if len(prev.Ops) != len(next.Ops) {
		return diffByPosition(prev, next)
	}

	dirty := image.Rectangle{}
	changed := 0

	for i := range prev.Ops {
		if sameOp(&prev.Ops[i], &next.Ops[i]) {
			continue
		}

		dirty = dirty.
			Union(opBounds(&prev.Ops[i], prev)).
			Union(opBounds(&next.Ops[i], next))
		changed++
	}

	return capDirty(next, dirty, changed)
}

// opKey identifies an operation by its paint kind and geometry.
type opKey struct {
	kind       layout.DisplayKind
	x, y, w, h float64
}

func keyOf(op *layout.DisplayOp) opKey {
	return opKey{op.Kind, op.X, op.Y, op.W, op.H}
}

// capDirty falls back to the frame when the change is too big or too spread
// out to merge.
func capDirty(next *layout.Display, dirty image.Rectangle, changed int) (image.Rectangle, bool) {
	if changed == 0 {
		return image.Rectangle{}, false
	}

	frame := image.Rect(0, 0, next.Width, next.Height)
	if changed > maxDirtyOps || dirty.Dx()*dirty.Dy() > frame.Dx()*frame.Dy()/3 {
		return frame, true
	}

	return dirty, false
}
