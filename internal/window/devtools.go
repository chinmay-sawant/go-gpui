package window

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// devState is the window's side of the devtools overlay: the mode flags, the
// boxes under the cursor and pinned, the panel rects, and the frame numbers.
// The page owns the on/off flag, so Config.DevTools and SetDevTools work.
type devState struct {
	on        bool
	ops       bool
	tab       devTab
	dockW     float64
	scroll    float64
	collapsed map[string]bool
	rows      []devRow
	content   devRect
	hovered   layout.Box
	haveHov   bool
	pinned    layout.Box
	havePin   bool
	opPick    int
	haveOp    bool
	resizing  bool
	gen       uint64
	stats     host.Stats
	frame     time.Duration
	draw      time.Duration
	lastAt    time.Time
	watch     keyWatch
	eaten     [len(devKeys)]bool
	panel     devRect
	hits      []devHit
	forward   bool
	// stroke is the outline hook. Nil strokes with vector.StrokeRect; tests
	// replace it to read the computed screen rects without pixels.
	stroke func(screen *ebiten.Image, r devRect, ink color.RGBA)
}

// devInspector returns the screen's inspector, or nil for a plain screen.
func (s *shell) devInspector() host.Inspector {
	insp, _ := s.app.(host.Inspector)

	return insp
}

// devSync follows the page's flag once per frame, so Config.DevTools and a
// SetDevTools call open or close the overlay without a key.
func (s *shell) devSync() error {
	insp := s.devInspector()
	if insp == nil {
		s.dev.on = false

		return nil
	}

	on := insp.DevTools()
	if on == s.dev.on {
		return nil
	}

	s.dev.on = on
	if !on {
		return s.devClose()
	}

	return nil
}
