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

// Clamp pulls a size into the min and max this page will draw.
func (p *Page) Clamp(width, height int) (int, int) {
	return clamp(width, p.minWidth, p.maxWidth), clamp(height, p.minHeight, p.maxHeight)
}

// SetSize stores the frame size used by the next Redraw.
// Values outside the min and max are pulled back inside that range.
func (p *Page) SetSize(width, height int) {
	p.width, p.height = p.Clamp(width, height)
}

// Generation increases by one on every successful Redraw.
func (p *Page) Generation() uint64 {
	return p.generation
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}
