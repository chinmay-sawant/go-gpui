package window

import "github.com/chinmay-sawant/go-gpui/internal/host"

// syncScrollWindow reports offset changes to a virtualized page. The shell
// keeps owning the offset; the page slices its rows and redraws only when
// it opted into windowing. Every other page scrolls by blitting, as before.
func (s *shell) syncScrollWindow() error {
	observer, ok := s.app.(host.ScrollObserver)
	if !ok {
		return nil
	}

	x, y := observer.ScrollOffset()
	if x == s.scrollX && y == s.scrollY {
		return nil
	}

	observer.SetScrollOffset(s.scrollX, s.scrollY)
	if !observer.Windowing() {
		return nil
	}

	return s.app.Redraw(s.ctx)
}
