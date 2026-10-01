package window

import "github.com/hajimehoshi/ebiten/v2"
import "github.com/hajimehoshi/ebiten/v2/inpututil"

func (s *shell) pointer() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if err := s.click(x, y); err != nil {
			return err
		}
	}

	for _, id := range inpututil.JustPressedTouchIDs() {
		x, y := ebiten.TouchPosition(id)
		if err := s.click(x, y); err != nil {
			return err
		}
	}

	return nil
}

func (s *shell) click(x, y int) error {
	frameW, frameH := s.frameSize()
	px, py := contentPoint(x, y, s.scrollX, s.scrollY, s.stretched(), frameW, frameH, s.screenW, s.screenH)

	return s.app.Click(s.ctx, px, py)
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
	wheelX, wheelY := ebiten.Wheel()
	frameW, frameH := s.frameSize()
	s.scrollX, s.scrollY = panScroll(
		s.scrollX, s.scrollY, wheelX, wheelY, frameW, frameH, s.screenW, s.screenH,
	)
}

func (s *shell) frameSize() (int, int) {
	if s.img != nil {
		bounds := s.img.Bounds()

		return bounds.Dx(), bounds.Dy()
	}

	return s.app.Size()
}
