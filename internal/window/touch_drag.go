package window

func (s *shell) routeTouchDrag(u touchUpdate, now []touchPos, w, h int) (bool, error) {
	if s.gesture.active && s.gesture.mouse {
		return false, nil
	}
	if !s.gesture.active && u.start != nil && len(now) == 1 {
		x, y := s.contentAt(u.start.x, u.start.y, w, h)
		if _, err := s.beginPointerDrag(x, y, false); err != nil {
			return false, err
		}
	}
	if !s.gesture.active {
		return false, nil
	}
	if len(now) == 0 {
		tap := u.tap != nil && !s.gesture.moved
		if err := s.endPointerDrag(false); err != nil {
			return true, err
		}
		return !tap, nil
	}
	if len(now) > 1 {
		return true, s.endPointerDrag(false)
	}
	x, y := s.contentAt(now[0].x, now[0].y, w, h)
	return true, s.movePointerDrag(x, y)
}
