package page

import "github.com/chinmay-sawant/ownframe/internal/host"

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
