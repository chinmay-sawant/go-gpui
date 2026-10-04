package page

import (
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// diffByPosition pairs operations by kind and geometry when the two lists
// differ in length. An operation with no partner dirties its bounds.
func diffByPosition(prev, next *layout.Display) (image.Rectangle, bool) {
	queue := map[opKey][]int{}
	for i := range next.Ops {
		k := keyOf(&next.Ops[i])
		queue[k] = append(queue[k], i)
	}

	seen := make([]bool, len(next.Ops))
	dirty := image.Rectangle{}
	changed := 0

	for i := range prev.Ops {
		k := keyOf(&prev.Ops[i])
		list := queue[k]

		if len(list) == 0 {
			dirty = dirty.Union(opBounds(&prev.Ops[i], prev))
			changed++

			continue
		}

		j := list[0]
		queue[k] = list[1:]
		seen[j] = true

		if sameOp(&prev.Ops[i], &next.Ops[j]) {
			continue
		}

		dirty = dirty.
			Union(opBounds(&prev.Ops[i], prev)).
			Union(opBounds(&next.Ops[j], next))
		changed++
	}

	for j := range next.Ops {
		if !seen[j] {
			dirty = dirty.Union(opBounds(&next.Ops[j], next))
			changed++
		}
	}

	return capDirty(next, dirty, changed)
}

// opBounds is one operation's box in CSS pixels. A text or bullet op carries
// its baseline in Y, so the top comes from the face ascent.
func opBounds(op *layout.DisplayOp, d *layout.Display) image.Rectangle {
	scale := d.PointsPerPixel
	if scale <= 0 {
		scale = 1
	}

	x0, y0 := op.X, op.Y
	x1, y1 := op.X+op.W, op.Y+op.H

	if op.Kind == layout.DisplayOpText || op.Kind == layout.DisplayOpBullet {
		y0 = op.Y - ascentOf(op)
		y1 = op.Y + op.InkDescent
	}

	return image.Rect(
		int(x0*scale)-1, int(y0*scale)-1,
		int(x1*scale)+1, int(y1*scale)+1,
	)
}

// ascentOf returns the face ascent in points for one text op.
func ascentOf(op *layout.DisplayOp) float64 {
	if op.Font == nil || op.Font.UnitsPerEm() <= 0 {
		return op.Size
	}

	return float64(op.Font.Ascent()) / float64(op.Font.UnitsPerEm()) * op.Size
}
