package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/replay"
)

// applyRepaint brings the persistent buffer up to date for this frame.
func (s *shell) applyRepaint(display *layout.Display, plan repaintPlan) {
	switch plan.mode {
	case repaintBuffer:
		s.ensureBuffer(plan.rect.Dx(), plan.rect.Dy())
		if s.partial.buf == nil {
			return
		}

		s.partial.buf.Fill(s.pageBackground())
		replay.DrawUnfixed(s.partial.buf, display, 0, 0)
	case repaintRect:
		if s.partial.buf == nil {
			return
		}

		if sub, ok := s.partial.buf.SubImage(plan.rect).(*ebiten.Image); ok {
			sub.Fill(s.pageBackground())
		}

		replay.DrawRect(s.partial.buf, display, plan.rect, 0, 0)
	}

	s.partial.gen = s.app.Generation()
}

// ensureBuffer keeps one content-sized image for the replay path. The content
// covers the canvas and any box that overflows it, so a scrolled page stays
// painted. The image is rebuilt when the content size changes.
func (s *shell) ensureBuffer(width, height int) {
	if s.partial.buf != nil && s.partial.width == width && s.partial.height == height {
		return
	}

	if s.partial.buf != nil {
		s.partial.buf.Dispose()
		s.partial.buf = nil
	}

	if width <= 0 || height <= 0 {
		return
	}

	s.partial.buf = ebiten.NewImage(width, height)
	s.partial.width, s.partial.height = width, height
}
