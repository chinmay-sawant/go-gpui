package window

import "github.com/hajimehoshi/ebiten/v2"

// resize lays the page out at the window size on the frame that size changes.
// The previous picture is not held and stretched while the drag continues.
func (s *shell) resize() error {
	wantW, wantH := s.app.Clamp(s.pendingW, s.pendingH)
	haveW, haveH := s.app.Size()
	if wantW == haveW && wantH == haveH {
		return nil
	}

	s.app.SetSize(wantW, wantH)

	return s.app.Redraw(s.ctx)
}

func (s *shell) syncImage() error {
	if s.app.Generation() == s.seq && s.img != nil {
		return nil
	}

	painted := s.app.Image()
	if painted == nil {
		return errNoImage
	}

	if s.img != nil {
		s.img.Dispose()
	}

	s.img = ebiten.NewImageFromImage(painted)
	s.seq = s.app.Generation()

	return nil
}

func (s *shell) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth < 1 {
		outsideWidth = 1
	}

	if outsideHeight < 1 {
		outsideHeight = 1
	}

	s.pendingW = outsideWidth
	s.pendingH = outsideHeight
	s.screenW = outsideWidth
	s.screenH = outsideHeight

	return outsideWidth, outsideHeight
}
