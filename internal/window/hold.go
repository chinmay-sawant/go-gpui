package window

import (
	"time"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// holdStart arms the held-press watch at a press point.
func (s *shell) holdStart(px, py float64) {
	s.hold.arm(time.Now(), px, py)
}

// holdFrame checks the held press from the frame loop, never a goroutine.
func (s *shell) holdFrame() error {
	return s.holdCheck(time.Now())
}

// holdCheck fires the screen's LongPress once when the press has been held
// in place long enough. A claim keeps the gesture for a selection drag.
func (s *shell) holdCheck(now time.Time) error {
	if !s.hold.due(now) {
		return nil
	}

	s.hold.fired = true

	lp, ok := s.app.(host.LongPresser)
	if !ok {
		return nil
	}

	claimed, err := lp.LongPress(s.ctx, s.hold.x, s.hold.y)
	if err != nil {
		return err
	}

	s.hold.claimed = claimed
	if claimed {
		s.dragActive = true
	}

	return nil
}
