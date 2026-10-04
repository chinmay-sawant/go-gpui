package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	devPanelWidth = 260
	devPanelPad   = 6
	devPanelGap   = 8
	devLineGap    = 2
)

// drawDevTools paints the overlay over the frame: the operation outlines,
// the hovered and pinned boxes, and the stats panel. It only reads shell
// state, so the page picture never changes.
func (s *shell) drawDevTools(screen *ebiten.Image) {
	lines := s.devLines()

	s.devLayout(lines)

	if s.dev.ops {
		s.drawDevOps(screen)
	}

	s.drawDevPick(screen)
	s.drawDevPanel(screen, lines)
}

// devLayout sizes the panel from the widest line and stores its clickable
// rows.
func (s *shell) devLayout(lines []string) {
	lineH := devLineHeight()
	w := float64(devPanelWidth)

	for _, line := range lines {
		if width, _ := text.Measure(line, badgeFace, 0); width+2*devPanelPad > w {
			w = width + 2*devPanelPad
		}
	}

	h := float64(len(lines))*lineH + 2*devPanelPad
	contentW, _ := s.contentSize()
	hScroll := !s.stretched() && scrollbarVisible(contentW, s.screenW)

	s.dev.panel = devPanelRect(s.screenW, s.screenH, w, h, hScroll)
	s.dev.hits = s.dev.hits[:0]

	row := devRect{
		X: s.dev.panel.X,
		Y: s.dev.panel.Y + devPanelPad + lineH,
		W: s.dev.panel.W,
		H: lineH,
	}
	s.dev.hits = append(s.dev.hits, devHit{rect: row, kind: devHitOps})
}
