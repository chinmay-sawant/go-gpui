package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// pointer sends hover and clicks. A click needs an up-to-down edge the shell
// saw itself. IsMouseButtonJustPressed can repeat for one press when a
// blocking call, such as the file dialog, stalls the event queue.
func (s *shell) pointer() error {
	x, y := ebiten.CursorPosition()
	if s.pointerScrollbar(x, y) {
		return nil
	}

	frameW, frameH := s.frameSize()
	px, py := contentPoint(x, y, s.scrollX, s.scrollY, s.stretched(), frameW, frameH, s.screenW, s.screenH)

	if err := s.app.Hover(s.ctx, px, py); err != nil {
		return err
	}

	down := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if down && s.mouseDown {
		s.dragScroll(y)
	}

	clicked := false
	if pressedNow(down, s.mouseDown) {
		clicked = true
		if err := s.pressAt(px, py); err != nil {
			return err
		}
	} else if down {
		if err := s.dragAt(px, py); err != nil {
			return err
		}
	}
	s.mouseDown = down

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if err := s.releaseAt(); err != nil {
			return err
		}
	}

	return s.touch(clicked, frameW, frameH)
}

// pressedNow reports the up-to-down edge of a pointer level.
func pressedNow(down, wasDown bool) bool {
	return down && !wasDown
}
