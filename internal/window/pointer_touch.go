package window

import "github.com/hajimehoshi/ebiten/v2"

// touch taps each fresh touch. Only one pointer action runs per frame, so a
// tap that the system also reports as a mouse click does not tap twice.
func (s *shell) touch(clicked bool, frameW, frameH int) error {
	now := ebiten.TouchIDs()
	fresh := freshTouches(now, s.touches)
	s.touches = append(s.touches[:0], now...)

	if clicked {
		return nil
	}

	for _, id := range fresh {
		tx, ty := ebiten.TouchPosition(id)
		tpx, tpy := contentPoint(tx, ty, s.scrollX, s.scrollY, s.stretched(), frameW, frameH, s.screenW, s.screenH)

		if err := s.app.Press(s.ctx, tpx, tpy); err != nil {
			return err
		}

		if err := s.app.Click(s.ctx, tpx, tpy); err != nil {
			return err
		}

		if err := s.app.Release(s.ctx); err != nil {
			return err
		}
	}

	return nil
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
