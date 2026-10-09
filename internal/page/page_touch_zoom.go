package page

// SetTouchZoomAllowed controls whether two-finger touch gestures zoom the page.
// It does not change scrolling, wheel zoom, keyboard zoom, or fitting.
func (p *Page) SetTouchZoomAllowed(on bool) {
	if p != nil {
		p.allowTouchZoom = on
	}
}

// TouchZoomAllowed reports whether two-finger touch gestures may zoom the page.
func (p *Page) TouchZoomAllowed() bool {
	return p == nil || p.allowTouchZoom
}
