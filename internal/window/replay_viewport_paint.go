package window

import (
	"image"

	"github.com/chinmay-sawant/blinkless/layout"
	"github.com/chinmay-sawant/ownframe/internal/replay"
	"github.com/hajimehoshi/ebiten/v2"
)

const viewportOverscan = 256

func (s *shell) setViewportArea(bounds image.Rectangle) {
	gap := min(viewportOverscan, bounds.Dy())
	width, height := bounds.Dx(), bounds.Dy()+2*gap
	if limit := ebiten.MaxImageSize(); limit > 0 {
		height = min(height, limit)
	}
	if s.viewport.buf == nil || s.viewport.buf.Bounds().Size() != image.Pt(width, height) {
		s.disposeViewport()
		s.viewport.buf = ebiten.NewImage(width, height)
	}
	top := max(0, s.scrollY-gap)
	s.viewport.area = image.Rect(s.scrollX, top, s.scrollX+width, top+height)
}

func (s *shell) paintViewport(d *layout.Display, rect image.Rectangle, pinZ int) {
	rect = rect.Intersect(s.viewport.area).Sub(s.viewport.area.Min)
	if rect.Empty() {
		return
	}
	clip := s.viewport.buf.SubImage(rect).(*ebiten.Image)
	clip.Clear()
	clip.Fill(s.pageBackground())
	replay.DrawVisibleUnfixedZ(clip, d, -float64(s.viewport.area.Min.X), -float64(s.viewport.area.Min.Y), pinZ)
}
