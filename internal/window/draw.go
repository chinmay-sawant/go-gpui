package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/replay"
)

func (s *shell) Draw(screen *ebiten.Image) {
	screen.Fill(s.pageBackground())

	if s.fallback {
		defer s.drawBadge(screen)
	}

	s.drawContent(screen)
	s.drawScrollbars(screen)
	s.drawMenu(screen)
}

func (s *shell) drawContent(screen *ebiten.Image) {
	zoom := s.zoom()

	if s.display != nil {
		if s.stretched() || zoom != 1 {
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
		if zoom != 1 {
			op.Filter = ebiten.FilterLinear
			op.GeoM.Scale(zoom, zoom)
		}

		op.GeoM.Translate(-float64(s.scrollX), -float64(s.scrollY))
		screen.DrawImage(s.img, &op)

		return
	}

	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(
		float64(s.screenW)/float64(bounds.Dx())*zoom,
		float64(s.screenH)/float64(bounds.Dy())*zoom,
	)
	screen.DrawImage(s.img, &op)
}
