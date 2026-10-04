package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// syncImage follows the page after each Redraw. A display-list page keeps
// its operations and drops any bitmap; any other page keeps one Ebiten image
// built from its picture. The old image is disposed when it is replaced.
func (s *shell) syncImage() error {
	if s.built() {
		return nil
	}

	if display := s.app.Display(); display != nil {
		s.setDisplay(display)
		s.pullScroll()

		return nil
	}

	painted := s.app.Image()
	if painted == nil {
		return errNoImage
	}

	next := ebiten.NewImageFromImage(painted)
	s.disposeImage()

	s.img = next
	s.display = nil
	s.fallback = true
	s.seq = s.app.Generation()
	s.pullScroll()

	return nil
}

// built reports that the cached artifact came from the page's current
// generation at the size the page committed, so no rebuild is needed.
func (s *shell) built() bool {
	if s.app.Generation() != s.seq || (s.img == nil && s.display == nil) {
		return false
	}

	w, h := s.app.Size()
	cw, ch := s.canvasSize()

	return cw == w && ch == h
}

func (s *shell) setDisplay(display *layout.Display) {
	s.disposeImage()

	s.display = display
	s.img = nil
	s.fallback = false
	s.seq = s.app.Generation()
}

func (s *shell) disposeImage() {
	if s.img != nil {
		s.img.Dispose()
		s.img = nil
	}
}
