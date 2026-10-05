package window

import (
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
	"github.com/hajimehoshi/ebiten/v2"
)

type viewportKey struct {
	generation uint64
	bounds     image.Rectangle
}

type viewportState struct {
	buf   *ebiten.Image
	key   viewportKey
	area  image.Rectangle
	valid bool
}

func (s *shell) drawViewport(dst *ebiten.Image, d *layout.Display, taker dirtyTaker) {
	s.dropContentBuffers()
	bounds := dst.Bounds()
	if bounds.Empty() {
		return
	}
	key := viewportKey{generation: s.seq, bounds: bounds}
	dirty, changed := taker.TakeDirty()
	wanted := bounds.Add(image.Pt(s.scrollX, s.scrollY))
	full := !s.viewport.valid || s.viewport.key.bounds != bounds || !wanted.In(s.viewport.area)
	full = full || (s.viewport.key.generation != key.generation && !changed)
	if full {
		s.setViewportArea(bounds)
		s.paintViewport(d, s.viewport.area)
	} else if changed {
		s.paintViewport(d, dirty)
	}
	s.viewport.key, s.viewport.valid = key, true
	var options ebiten.DrawImageOptions
	options.GeoM.Translate(float64(s.viewport.area.Min.X-s.scrollX), float64(s.viewport.area.Min.Y-s.scrollY))
	dst.DrawImage(s.viewport.buf, &options)
}

func (s *shell) disposeViewport() {
	if s.viewport.buf != nil {
		s.viewport.buf.Dispose()
	}
	s.viewport = viewportState{}
}

func (s *shell) dropContentBuffers() {
	if s.partial.buf != nil {
		s.partial.buf.Dispose()
		s.partial = partialState{}
	}
	if s.replayBuf != nil {
		s.replayBuf.Dispose()
		s.replayBuf = nil
	}
}
