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

// dirtyFromDisplay adds the difference between prev and display to the
// dirty region. The first draw and a size change dirty the frame. With Perf
// on it also stores the changed, dirty, and region counters Stats reports;
// TakeDirty clears only the region, never these counters.
func (p *Page) dirtyFromDisplay(prev, display *Display) {
	if p.dirtyFull {
		return
	}

	if prev == nil {
		p.markFull()
		if p.perf {
			p.recordFull(len(display.Ops))
		}

		return
	}

	r, full, changed := diffDisplay(prev, display)
	if full {
		p.markFull()
		if p.perf {
			p.recordFull(len(display.Ops))
		}

		return
	}

	p.markRect(r)
	if p.perf {
		p.recordPartial(changed, r, display)
	}
}

// applyPending unions the new box of every pending id, so an element that
// grew or shrank repaints the space it now uses too.
func (p *Page) applyPending() {
	for id := range p.pending {
		p.markElement(id)
	}

	p.pending = nil
}
