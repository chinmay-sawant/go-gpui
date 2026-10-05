package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/replay"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// oversized reports a content size no Ebiten image can hold. A zero device
// limit means the game is not running, so nothing is oversized there: no
// GPU image exists to overflow.
func oversized(w, h int) bool {
	m := ebiten.MaxImageSize()
	if m <= 0 {
		return false
	}

	return w > m || h > m
}

// isTicking reports a screen that changes operations from a frame callback.
// A tick changes operations in place without a dirty rect, so such a page
// keeps the full replay.
func (s *shell) isTicking() bool {
	t, ok := s.app.(ticking)

	return ok && t.Ticking()
}

// directReplay draws the list straight to the screen. Ticking pages use it
// because they have no dirty rect; oversized content uses it because no
// buffer can hold it.
func (s *shell) directReplay(dst *ebiten.Image, display *layout.Display) {
	replay.Draw(dst, display, -float64(s.scrollX), -float64(s.scrollY))
}
