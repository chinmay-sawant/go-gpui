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
		s.ensureBuffer(display)
		if s.partial.buf == nil {
			return
		}

		s.partial.buf.Fill(s.pageBackground())
		replay.Draw(s.partial.buf, display, 0, 0)
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

// ensureBuffer keeps one canvas-sized image for the replay path. The image is
// rebuilt when the page size changes.
func (s *shell) ensureBuffer(display *layout.Display) {
	if s.partial.buf != nil && s.partial.width == display.Width && s.partial.height == display.Height {
		return
	}

	if s.partial.buf != nil {
		s.partial.buf.Dispose()
		s.partial.buf = nil
	}

	if display.Width <= 0 || display.Height <= 0 {
		return
	}

	s.partial.buf = ebiten.NewImage(display.Width, display.Height)
	s.partial.width, s.partial.height = display.Width, display.Height
}
