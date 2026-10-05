package window

import "github.com/hajimehoshi/ebiten/v2"

// touch feeds the live fingers to the gesture tracker. A finger that lifts
// without moving taps; a moved finger drags the page; two fingers pinch.
// Only one pointer action runs per frame, so a tap that the system also
// reports as a mouse click does not tap twice.
func (s *shell) touch(clicked bool, frameW, frameH int) error {
	ids := ebiten.TouchIDs()
	now := make([]touchPos, 0, len(ids))

	for _, id := range ids {
		x, y := ebiten.TouchPosition(id)
		now = append(now, touchPos{id: id, x: x, y: y})
	}

	u := s.fingers.frame(now, clicked)

	if (u.dx != 0 || u.dy != 0) && !s.stretched() {
		contentW, contentH := s.contentSize()
		s.scrollX, s.scrollY = clampScroll(
			s.scrollX-u.dx, s.scrollY-u.dy, contentW, contentH, s.screenW, s.screenH,
		)
	}

	if u.tap == nil {
		return nil
	}

	tpx, tpy := s.contentAt(u.tap.x, u.tap.y, frameW, frameH)

	return s.tapAt(tpx, tpy)
}

// tapAt forwards a tap as press, caret, click, release. The caret lands
// before the click handler runs, matching pressAt. A page without the
// selector keeps the old press, click, release behavior.
func (s *shell) tapAt(px, py float64) error {
	if err := s.app.Press(s.ctx, px, py); err != nil {
		return err
	}

	if sel, ok := s.app.(selector); ok {
		if err := sel.SelectAt(s.ctx, px, py); err != nil {
			return err
		}
	}

	if err := s.app.Click(s.ctx, px, py); err != nil {
		return err
	}

	return s.app.Release(s.ctx)
}

// freshTouches returns the ids in now that are not in prev.
func freshTouches(now, prev []ebiten.TouchID) []ebiten.TouchID {
	var fresh []ebiten.TouchID

	for _, id := range now {
		if !hasTouch(prev, id) {
			fresh = append(fresh, id)
		}
	}

	return fresh
}

func hasTouch(ids []ebiten.TouchID, want ebiten.TouchID) bool {
	for _, id := range ids {
		if id == want {
			return true
		}
	}

	return false
}
