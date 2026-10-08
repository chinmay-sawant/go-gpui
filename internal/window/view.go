package window

import "github.com/hajimehoshi/ebiten/v2"

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
	if s.devWheel() {
		return
	}

	wheelX, wheelY := ebiten.Wheel()

	if s.zoomWheel(wheelY) {
		return
	}

	if s.stretched() || !s.allowPageScroll() {
		return
	}

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
