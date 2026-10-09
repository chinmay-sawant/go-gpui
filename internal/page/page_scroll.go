package page

import "github.com/chinmay-sawant/ownframe/internal/host"

// SetAllowScroll turns window panning on or off.
// The pad screen turns it off so a drag moves the TV pointer.
func (p *Page) SetAllowScroll(on bool) {
	if p == nil {
		return
	}

	p.allowScroll = on
}

// AllowScroll reports that the window may pan this page.
func (p *Page) AllowScroll() bool {
	if p == nil {
		return true
	}

	return p.allowScroll
}

// ScrollTo queues an absolute scroll move for the window. The window clamps
// the target to the content size and owns the offset.
func (p *Page) ScrollTo(x, y int) {
	p.scroll = host.Scroll{X: x, Y: y, Absolute: true}
	p.hasScroll = true
}

// ScrollBy queues a scroll delta for the window. A second delta before the
// window consumes it adds to the first, and a delta after ScrollTo moves the
// pending target.
func (p *Page) ScrollBy(dx, dy int) {
	p.scroll.X += dx
	p.scroll.Y += dy
	p.hasScroll = true
}

// TakeScroll returns the queued request and clears it.
func (p *Page) TakeScroll() (host.Scroll, bool) {
	req, ok := p.scroll, p.hasScroll
	p.scroll, p.hasScroll = host.Scroll{}, false

	return req, ok
}
