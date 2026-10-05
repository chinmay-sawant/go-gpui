package page

import "image"

// TakeDirty returns the region in CSS pixels changed since the previous
// call and clears it. A whole-content change returns the content rect. ok
// is false when nothing changed since the previous call.
func (p *Page) TakeDirty() (image.Rectangle, bool) {
	full := p.dirtyFull
	r := p.dirty
	p.dirtyFull = false
	p.dirty = image.Rectangle{}
	p.pending = nil

	content := p.contentRect()
	if full {
		return content, true
	}

	if r.Empty() {
		return image.Rectangle{}, false
	}

	r = r.Inset(-dirtyPad).Intersect(content)
	if r.Empty() {
		r = content
	}

	return r, true
}

// contentRect is the painted page in CSS pixels: the canvas plus any box
// that overflows it, the region the window's replay buffer covers.
func (p *Page) contentRect() image.Rectangle {
	w, h := p.width, p.height
	if p.display != nil {
		w, h = p.display.Width, p.display.Height
	}

	r := image.Rect(0, 0, w, h)
	for _, b := range p.boxes {
		r = r.Union(boxRect(b))
	}

	return r
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
