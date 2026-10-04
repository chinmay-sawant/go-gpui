package page

import "image"

// TakeDirty returns the region in CSS pixels changed since the previous
// call and clears it. A whole-frame change returns the frame rect. ok is
// false when nothing changed since the previous call.
func (p *Page) TakeDirty() (image.Rectangle, bool) {
	full := p.dirtyFull
	r := p.dirty
	p.dirtyFull = false
	p.dirty = image.Rectangle{}
	p.pending = nil

	frame := image.Rect(0, 0, p.width, p.height)
	if full {
		return frame, true
	}

	if r.Empty() {
		return image.Rectangle{}, false
	}

	r = r.Inset(-dirtyPad).Intersect(frame)
	if r.Empty() {
		r = frame
	}

	return r, true
}

// dirtyFromDisplay adds the difference between prev and display to the
// dirty region. The first draw and a size change dirty the frame.
func (p *Page) dirtyFromDisplay(prev, display *Display) {
	if p.dirtyFull {
		return
	}

	if prev == nil {
		p.markFull()

		return
	}

	r, full := diffDisplay(prev, display)
	if full {
		p.markFull()

		return
	}

	p.markRect(r)
}

// applyPending unions the new box of every pending id, so an element that
// grew or shrank repaints the space it now uses too.
func (p *Page) applyPending() {
	for id := range p.pending {
		p.markElement(id)
	}

	p.pending = nil
}
