package window

import "github.com/chinmay-sawant/ownframe/internal/host"

func (s *shell) applyTouches(now []touchPos, clicked bool, frameW, frameH int) error {
	lifted := len(s.fingers.fingers) > 0 && len(now) == 0
	u := s.fingers.frame(now, clicked)
	if s.viewLocked() {
		s.holdView()
	} else if !s.touchZoomAllowed() {
		s.holdTouchZoom()
	}

	if handled, err := s.routeTouchDrag(u, now, frameW, frameH); handled || err != nil {
		return err
	}

	if len(now) >= 2 {
		s.hold.cancel()
	} else if u.start != nil {
		tpx, tpy := s.contentAt(u.start.x, u.start.y, frameW, frameH)
		s.holdStart(tpx, tpy)
		if err := s.app.Press(s.ctx, tpx, tpy); err != nil {
			return err
		}
	}

	if err := s.touchMove(u, now, frameW, frameH); err != nil {
		return err
	}

	if u.swipe != nil && !s.hold.claimed {
		if sw, ok := s.app.(host.Swiper); ok {
			if err := sw.Swipe(s.ctx, float64(u.swipe.dx), float64(u.swipe.dy)); err != nil {
				return err
			}
		}
	}

	if u.tap == nil {
		if lifted {
			return s.app.Release(s.ctx)
		}
		return nil
	}

	if s.hold.claimed {
		return s.releaseAt()
	}

	tpx, tpy := s.contentAt(u.tap.x, u.tap.y, frameW, frameH)

	return s.tapAt(tpx, tpy)
}
