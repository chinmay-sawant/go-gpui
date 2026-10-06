package window

import "github.com/hajimehoshi/ebiten/v2"

// touch feeds the live fingers to the gesture tracker. A finger that lifts
// without moving taps; a moved finger drags the page; two fingers pinch.
// A held finger arms the long press. Only one pointer action runs per
// frame, so a tap that the system also reports as a mouse click does not
// tap twice.
func (s *shell) touch(clicked bool, frameW, frameH int) error {
	ids := ebiten.TouchIDs()
	now := make([]touchPos, 0, len(ids))

	for _, id := range ids {
		x, y := ebiten.TouchPosition(id)
		now = append(now, touchPos{id: id, x: x, y: y})
	}

	u := s.fingers.frame(now, clicked)

	if len(now) >= 2 {
		s.hold.cancel()
	} else if u.start != nil {
		tpx, tpy := s.contentAt(u.start.x, u.start.y, frameW, frameH)
		s.holdStart(tpx, tpy)
	}

	if err := s.touchMove(u, now, frameW, frameH); err != nil {
		return err
	}

	if u.tap == nil {
		return nil
	}

	if s.hold.claimed {
		return s.releaseAt()
	}

	tpx, tpy := s.contentAt(u.tap.x, u.tap.y, frameW, frameH)

	return s.tapAt(tpx, tpy)
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
