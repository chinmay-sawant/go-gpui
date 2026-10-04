package window

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/replay"
)

func (s *shell) Draw(screen *ebiten.Image) {
	s.devFrameTime()

	screen.Fill(s.pageBackground())

	if s.fallback {
		defer s.drawBadge(screen)
	}

	start := time.Now()
	s.drawContent(screen)
	s.dev.draw = time.Since(start)

	s.drawScrollbars(screen)

	if s.dev.on {
		s.drawDevTools(screen)
	}
}

func (s *shell) drawContent(screen *ebiten.Image) {
	if s.display != nil {
		if s.stretched() {
			s.drawReplayScaled(screen)

			return
		}

		replay.Draw(screen, s.display, -float64(s.scrollX), -float64(s.scrollY))

		return
	}

	if s.img == nil {
		return
	}

	bounds := s.img.Bounds()
	if bounds.Dx() == 0 || bounds.Dy() == 0 || s.screenW == 0 || s.screenH == 0 {
		return
	}

	var op ebiten.DrawImageOptions
	// A page larger than the window keeps its own pixels and scrolls.
	// A size that has not been laid out yet is scaled until the new picture arrives.
	if !s.stretched() {
		op.GeoM.Translate(-float64(s.scrollX), -float64(s.scrollY))
		screen.DrawImage(s.img, &op)

		return
	}

	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(
		float64(s.screenW)/float64(bounds.Dx()),
		float64(s.screenH)/float64(bounds.Dy()),
	)
	screen.DrawImage(s.img, &op)
}
