package player

import (
	"math"
	"time"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/internal/frame"
)

// eqHeight is bar i's height as a fraction of its container. A paused
// player keeps every bar at a small static height.
func eqHeight(i int, playing bool, t time.Time) float64 {
	if !playing {
		return 0.3
	}

	f := 0.5 + 0.5*math.Sin(float64(t.UnixMilli())/float64(120+35*i)+float64(i)*1.9)

	return 0.2 + 0.8*f
}

// equalizer waves the bars in the queue head and the active row.
func (a *App) equalizer(d *gpui.Display, playing bool) {
	now := time.Now()

	for _, id := range []string{"eq", "eq-row"} {
		box, ok := boxByID(a.page.Boxes(), id)
		if !ok {
			continue
		}

		_, y, _, h := frame.BoxUnits(d, box)

		for i, bar := range frame.Fills(d, box, accent) {
			bar.H = h * eqHeight(i, playing, now)
			bar.Y = y + h - bar.H
		}
	}
}
