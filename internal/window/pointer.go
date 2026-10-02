package window

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/hajimehoshi/ebiten/v2/inpututil"

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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if err := s.app.Press(s.ctx, px, py); err != nil {
			return err
		}

		if err := s.app.Click(s.ctx, px, py); err != nil {
			return err
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if err := s.app.Release(s.ctx); err != nil {
			return err
		}
	}

	for _, id := range inpututil.JustPressedTouchIDs() {
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

// stretched reports that the painted page is being scaled to the window.
// A page larger than a window of the same layout size is scrolled instead.
func (s *shell) stretched() bool {
	frameW, frameH := s.frameSize()
	if frameW == s.screenW && frameH == s.screenH {
		return false
	}

	haveW, haveH := s.app.Size()
	if haveW == s.screenW && haveH == s.screenH && (frameW > s.screenW || frameH > s.screenH) {
		return false
	}

	return true
}

func (s *shell) wheel() {
	if s.stretched() {
		return
	}

	wheelX, wheelY := ebiten.Wheel()
	contentW, contentH := s.contentSize()
	s.scrollX, s.scrollY = panScroll(
		s.scrollX, s.scrollY, wheelX, wheelY, contentW, contentH, s.screenW, s.screenH,
	)
}

func (s *shell) frameSize() (int, int) {
	if s.img != nil {
		bounds := s.img.Bounds()

		return bounds.Dx(), bounds.Dy()
	}

	if s.display != nil {
		return s.display.Width, s.display.Height
	}

	return s.app.Size()
}
