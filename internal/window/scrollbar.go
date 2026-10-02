package window

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	scrollbarThickness = 8
	scrollbarMinThumb  = 24
)

var (
	scrollTrack = color.RGBA{A: 32}
	scrollThumb = color.RGBA{A: 128}
)

// contentSize returns the painted page size.
func (s *shell) contentSize() (int, int) {
	if s.display != nil {
		return s.display.Width, s.display.Height
	}

	if s.img != nil {
		b := s.img.Bounds()

		return b.Dx(), b.Dy()
	}

	return s.app.Size()
}

// scrollbarVisible reports whether content overflows the viewport.
func scrollbarVisible(content, viewport int) bool {
	return content > viewport && viewport > 0
}

// scrollbarThumb returns the thumb position and length along a track.
func scrollbarThumb(track, content, viewport, offset int) (float32, float32) {
	length := float32(track) * float32(viewport) / float32(content)
	if length < scrollbarMinThumb {
		length = scrollbarMinThumb
	}

	if length > float32(track) {
		length = float32(track)
	}

	maxOffset := content - viewport
	pos := float32(0)

	if maxOffset > 0 && float32(track) > length {
		pos = (float32(track) - length) * float32(offset) / float32(maxOffset)
	}

	return pos, length
}

// scrollbarOffset maps a thumb position back to a scroll offset.
func scrollbarOffset(pos float64, track, content, viewport int) int {
	_, length := scrollbarThumb(track, content, viewport, 0)
	span := float64(track) - float64(length)
	maxOffset := content - viewport

	if maxOffset <= 0 || span <= 0 {
		return 0
	}

	if pos < 0 {
		pos = 0
	}

	if pos > span {
		pos = span
	}

	return int(pos * float64(maxOffset) / span)
}

// drawScrollbars overlays the bars a page larger than the window needs.
func (s *shell) drawScrollbars(dst *ebiten.Image) {
	if s.stretched() {
		return
	}

	contentW, contentH := s.contentSize()
	thick := float32(scrollbarThickness)

	if scrollbarVisible(contentH, s.screenH) {
		pos, length := scrollbarThumb(s.screenH, contentH, s.screenH, s.scrollY)
		x := float32(s.screenW) - thick
		vector.FillRect(dst, x, 0, thick, float32(s.screenH), scrollTrack, false)
		vector.FillRect(dst, x, pos, thick, length, scrollThumb, false)
	}

	if scrollbarVisible(contentW, s.screenW) {
		pos, length := scrollbarThumb(s.screenW, contentW, s.screenW, s.scrollX)
		y := float32(s.screenH) - thick
		vector.FillRect(dst, 0, y, float32(s.screenW), thick, scrollTrack, false)
		vector.FillRect(dst, pos, y, length, thick, scrollThumb, false)
	}
}
