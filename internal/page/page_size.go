package page

// Handle registers the input functions. A later call replaces them.
func (p *Page) Handle(h Handlers) {
	p.handlers = h
}

// SetData stores the value the template prints on the next Redraw.
func (p *Page) SetData(data any) {
	p.data = data
	p.invalidateCache()
}

// Title returns the window title.
func (p *Page) Title() string {
	return p.title
}

// Size returns the frame size in CSS pixels.
func (p *Page) Size() (int, int) {
	return p.width, p.height
}

// MinSize returns the smallest frame this page will draw.
func (p *Page) MinSize() (int, int) {
	return p.minWidth, p.minHeight
}

// MaxSize returns the largest window the page asks its host for. The
// window clamps to it; the layout does not.
func (p *Page) MaxSize() (int, int) {
	return p.maxWidth, p.maxHeight
}

// Clamp pulls a size up to the minimum only, so the layout always tracks
// the window. The maximum is a window bound, not a layout bound.
func (p *Page) Clamp(width, height int) (int, int) {
	return clampMin(width, p.minWidth), clampMin(height, p.minHeight)
}

// SetSize stores the frame size used by the next Redraw.
// Values below the minimum are pulled up to it.
func (p *Page) SetSize(width, height int) {
	p.width, p.height = p.Clamp(width, height)
}

// Generation increases by one on every successful Redraw.
func (p *Page) Generation() uint64 {
	return p.generation
}

func clampMin(v, min int) int {
	if v < min {
		return min
	}

	return v
}
