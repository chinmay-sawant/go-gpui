package page

import "image"

// maxDirtyScan caps the ops countDirtyOps inspects, so one huge frame
// cannot make the dirty count cost more than the replay it measures.
const maxDirtyScan = 4096

// recordFull stores whole-frame dirty counters for a first draw, a size
// change, or a paint fallback. n is the ops or boxes now on screen.
func (p *Page) recordFull(n int) {
	p.stats.changedOps = n
	p.stats.dirtyOps = n
	p.stats.dirtyRegions = 1
}

// recordPartial stores the counters for one display diff. An empty rect
// means nothing changed; any other rect dirties one region.
func (p *Page) recordPartial(changed int, r image.Rectangle, d *Display) {
	p.stats.changedOps = changed
	if r.Empty() {
		p.stats.dirtyOps = 0
		p.stats.dirtyRegions = 0

		return
	}

	p.stats.dirtyOps = countDirtyOps(d, r)
	p.stats.dirtyRegions = 1
}

// countDirtyOps counts the ops whose bounds meet r, stopping at
// maxDirtyScan so the count stays cheaper than a full replay.
func countDirtyOps(d *Display, r image.Rectangle) int {
	if d == nil {
		return 0
	}

	n := 0

	for i := range d.Ops {
		if i >= maxDirtyScan {
			break
		}

		if opBounds(&d.Ops[i], d).Intersect(r).Empty() {
			continue
		}

		n++
	}

	return n
}
