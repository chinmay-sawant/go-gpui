package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// devPointer routes the pointer while the overlay is on. The panel owns its
// rects; the content pins the box under the cursor. Alt+click forwards the
// press and click to the page, and release follows a forwarded press.
func (s *shell) devPointer(x, y int, px, py float64) error {
	down := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	pressed := pressedNow(down, s.mouseDown)
	s.mouseDown = down
	released := inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)

	if devInRect(s.dev.panel, float64(x), float64(y)) {
		if pressed {
			return s.devPanelPress(x, y)
		}

		return nil
	}

	s.devHover(px, py)

	if pressed {
		return s.devPick(px, py, ebiten.IsKeyPressed(ebiten.KeyAlt))
	}

	if released && s.dev.forward {
		s.dev.forward = false

		return s.app.Release(s.ctx)
	}

	return nil
}

// devPanelPress hit-tests the panel's own rows. A panel click never reaches
// the page; the ops row toggles the view and any panel click clears the pin.
func (s *shell) devPanelPress(x, y int) error {
	s.dev.havePin = false

	for _, hit := range s.dev.hits {
		if devInRect(hit.rect, float64(x), float64(y)) && hit.kind == devHitOps {
			s.dev.ops = !s.dev.ops
		}
	}

	return nil
}

// devPick handles a press inside the content while the overlay is on. Alt
// forwards the press and click to the page; any other press pins the box
// under the cursor, and the same box again clears the pin.
func (s *shell) devPick(px, py float64, alt bool) error {
	if alt {
		s.dev.forward = true

		if err := s.app.Press(s.ctx, px, py); err != nil {
			return err
		}

		return s.app.Click(s.ctx, px, py)
	}

	box, ok := devBoxAt(s.app.Boxes(), px, py)
	if !ok {
		s.dev.havePin = false

		return nil
	}

	if s.dev.havePin && devSameBox(s.dev.pinned, box) {
		s.dev.havePin = false

		return nil
	}

	s.dev.pinned = box
	s.dev.havePin = true

	return nil
}
