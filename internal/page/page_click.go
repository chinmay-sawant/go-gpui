package page

import (
	"context"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// Click hit-tests the last picture and calls the click handler.
// The innermost box is the last one in document order that contains the point.
// A routed data-action loads that HTML and skips the handler.
// Click then draws the page again.
func (p *Page) Click(ctx context.Context, x, y float64) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	box, ok := hit(p.Boxes(), x, y)
	if ok && box.Action != "" {
		if html, routed := p.routes[box.Action]; routed {
			return p.Load(ctx, html)
		}
	}

	if ok && p.handlers.Click != nil {
		if err := p.handlers.Click(ctx, box); err != nil {
			return err
		}
	}

	return p.Redraw(ctx)
}

// hit returns the last box that contains x, y.
// Boxes are in document order, so that box is the inner element.
func hit(boxes []Box, x, y float64) (Box, bool) {
	var found Box
	ok := false

	for _, b := range boxes {
		if x < b.X || y < b.Y || x > b.X+b.W || y > b.Y+b.H {
			continue
		}

		found = b
		ok = true
	}

	return found, ok
}

var _ host.Screen = (*Page)(nil)
