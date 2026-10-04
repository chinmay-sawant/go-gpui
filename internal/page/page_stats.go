package page

import (
	"time"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// pageStats counts the stages of every Redraw. The devtools overlay reads
// them through Stats.
type pageStats struct {
	redraws    uint64
	parses     uint64
	cascades   uint64
	layouts    uint64
	repaints   uint64
	lastRedraw time.Duration
	lastDraw   time.Duration
	frameDraw  time.Duration
}

// Stats returns one snapshot of the page counters and the last frame sizes.
// LastDraw is the window's draw time once SetDrawTime has recorded one, and
// the display-list or paint stage of the last Redraw before that.
func (p *Page) Stats() host.Stats {
	ops := 0
	if p.display != nil {
		ops = len(p.display.Ops)
	}

	lastDraw := p.stats.lastDraw
	if p.stats.frameDraw > 0 {
		lastDraw = p.stats.frameDraw
	}

	return host.Stats{
		Redraws:    p.stats.redraws,
		Parses:     p.stats.parses,
		Cascades:   p.stats.cascades,
		Layouts:    p.stats.layouts,
		Repaints:   p.stats.repaints,
		Boxes:      len(p.boxes),
		Ops:        ops,
		LastRedraw: p.stats.lastRedraw,
		LastDraw:   lastDraw,
	}
}
