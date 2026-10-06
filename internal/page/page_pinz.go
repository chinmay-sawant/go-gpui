package page

// SetViewportPinZ pins the page's top layers. Operations whose z-index is at
// least z are drawn at the offset the display was built with, so they stay at
// the viewport while the page scrolls and need no redraw. Zero disables the
// pinning and draws every operation with the scroll, as before.
func (p *Page) SetViewportPinZ(z int) { p.pinZ = z }

// ViewportPinZ is the pinned layer threshold.
func (p *Page) ViewportPinZ() int { return p.pinZ }
