package window

import (
	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// applyScrollRequest consumes a page scroll request and moves the offset.
// The shell keeps ownership of the offset; the request is clamped to the
// content and cleared either way.
func (s *shell) applyScrollRequest() {
	requester, ok := s.app.(host.ScrollRequester)
	if !ok {
		return
	}

	req, ok := requester.TakeScroll()
	if !ok {
		return
	}

	contentW, contentH := s.contentSize()
	if req.Absolute {
		s.scrollX, s.scrollY = clampScroll(req.X, req.Y, contentW, contentH, s.screenW, s.screenH)

		return
	}

	s.scrollX, s.scrollY = clampScroll(
		s.scrollX+req.X, s.scrollY+req.Y, contentW, contentH, s.screenW, s.screenH,
	)
}
