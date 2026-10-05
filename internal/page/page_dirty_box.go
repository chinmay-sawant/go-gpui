package page

import "image"

// dirtyPad is the margin a dirty rect gains for shadows, outlines, and
// antialiasing.
const dirtyPad = 4

// labelGap is how near a label must sit to a toggle to repaint with it.
const labelGap = 12

// growBox covers the boxes inside r, so a nested span repaints with its
// parent.
func (p *Page) growBox(r image.Rectangle) image.Rectangle {
	if r.Empty() {
		return r
	}

	for _, b := range p.boxes {
		inner := boxRect(b)
		if !inner.Empty() && inner.In(r) {
			r = r.Union(inner)
		}
	}

	return r
}

// hasBox reports whether the last layout has a box for id.
func (p *Page) hasBox(id string) bool {
	for _, b := range p.boxes {
		if b.ID == id {
			return true
		}
	}

	return false
}

// boxRect is one box in whole CSS pixels.
func boxRect(b Box) image.Rectangle {
	return image.Rect(int(b.X), int(b.Y), int(b.X+b.W+0.5), int(b.Y+b.H+0.5))
}

// touches reports whether two rects meet within gap pixels.
func touches(a, b image.Rectangle, gap int) bool {
	return !a.Inset(-gap).Intersect(b).Empty()
}

// markRect grows the dirty region.
func (p *Page) markRect(r image.Rectangle) {
	if r.Empty() {
		return
	}

	p.dirty = p.dirty.Union(r)
}

// markFull dirties the whole content; a later TakeDirty returns the
// content rect.
func (p *Page) markFull() {
	p.dirtyFull = true
}
