package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/replay"
)

// drawReplayScaled draws the display list through an offscreen buffer and
// scales it to the window, matching the bitmap path when the page size and
// the window size differ (a resized or clamped window).
func (s *shell) drawReplayScaled(dst *ebiten.Image) {
	w, h := s.display.Width, s.display.Height
	if w <= 0 || h <= 0 || s.screenW <= 0 || s.screenH <= 0 {
		return
	}

	if s.replayBuf == nil || s.replayBuf.Bounds().Dx() != w || s.replayBuf.Bounds().Dy() != h {
		if s.replayBuf != nil {
			s.replayBuf.Dispose()
		}

		s.replayBuf = ebiten.NewImage(w, h)
	}

	s.replayBuf.Clear()
	s.replayBuf.Fill(s.pageBackground())
	replay.Draw(s.replayBuf, s.display, 0, 0)

	var op ebiten.DrawImageOptions
	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(float64(s.screenW)/float64(w), float64(s.screenH)/float64(h))
	dst.DrawImage(s.replayBuf, &op)
}
