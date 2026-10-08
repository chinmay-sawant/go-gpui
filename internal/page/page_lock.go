package page

// ViewLocked reports that the window should scale this page to the screen
// and ignore scroll and pinch zoom.
func (p *Page) ViewLocked() bool {
	if p == nil {
		return false
	}

	return p.lockView
}
