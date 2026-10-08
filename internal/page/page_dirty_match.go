package page

import (
	"image"

	"github.com/chinmay-sawant/blinkless/layout"
)

// diffByPosition pairs operations by kind and geometry when the two lists
// differ in length. An operation with no partner dirties its bounds, and
// changed counts every unpaired or repainted operation.
func diffByPosition(prev, next *layout.Display) (image.Rectangle, bool, int) {
	queue := make(map[opKey][]int, len(next.Ops))
	for i := range next.Ops {
		k := keyOf(&next.Ops[i])
		queue[k] = append(queue[k], i)
	}

	seen := make([]bool, len(next.Ops))
	dirty := image.Rectangle{}
	changed := 0
	frame := image.Rect(0, 0, next.Width, next.Height)

	for i := range prev.Ops {
		k := keyOf(&prev.Ops[i])
		list := queue[k]

		if len(list) == 0 {
			dirty = dirty.Union(opBounds(&prev.Ops[i], prev))
			changed++

			if changed > maxDirtyOps {
				return frame, true, len(next.Ops)
			}

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

		if changed > maxDirtyOps {
			return frame, true, len(next.Ops)
		}
	}

	for j := range next.Ops {
		if !seen[j] {
			dirty = dirty.Union(opBounds(&next.Ops[j], next))
			changed++

			if changed > maxDirtyOps {
				return frame, true, len(next.Ops)
			}
		}
	}

	r, full := capDirty(next, dirty, changed)

	return r, full, changed
}
