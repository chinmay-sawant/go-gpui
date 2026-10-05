package page

// StepScrollWindow prepares the visible rows and reports whether they changed.
// A windowing page without a callback keeps the ordinary scroll redraw path.
func (p *Page) StepScrollWindow() bool {
	if p.windowing.fn == nil {
		return true
	}
	return p.windowing.fn(p.windowing.y, p.height)
}
