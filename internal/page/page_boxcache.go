package page

import "image"

// setBoxes stores the hit-test boxes from a Redraw. It rebuilds the id
// index so lookups stay O(1) on large pages, and drops the cached content
// rect. Boxes change only here, so the index never goes stale.
func (p *Page) setBoxes(boxes []Box) {
	p.boxes = boxes
	p.index = map[string]Box{}
	for _, b := range boxes {
		if b.ID != "" {
			p.index[b.ID] = b
		}
	}
	p.contentOK = false
}

// boxByID returns the last box in p.boxes with that id, or Box{ID: id}.
func (p *Page) boxByID(id string) Box {
	if b, ok := p.index[id]; ok {
		return b
	}
	box := Box{ID: id}
	for _, b := range p.boxes {
		if b.ID == id {
			box = b
		}
	}
	return box
}

// hasBox reports whether the last layout has a box for id.
func (p *Page) hasBox(id string) bool {
	if _, ok := p.index[id]; ok {
		return true
	}
	for _, b := range p.boxes {
		if b.ID == id {
			return true
		}
	}
	return false
}

// contentRect is the painted page in CSS pixels: the canvas plus any box
// that overflows it, the region the window's replay buffer covers. The
// rect is cached per canvas size until the next setBoxes, so the per-frame
// TakeDirty costs O(1) on large pages.
func (p *Page) contentRect() image.Rectangle {
	w, h := p.width, p.height
	if p.display != nil {
		w, h = p.display.Width, p.display.Height
	}
	if p.contentOK && p.contentW == w && p.contentH == h {
		return p.content
	}
	r := image.Rect(0, 0, w, h)
	for _, b := range p.boxes {
		r = r.Union(boxRect(b))
	}
	p.content, p.contentW, p.contentH, p.contentOK = r, w, h, true
	return r
}
