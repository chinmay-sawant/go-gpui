package window

import (
	"math"
	"time"
)

const (
	clickGap  = 400 * time.Millisecond
	clickSlop = 4.0
)

// clickWatch counts rapid presses at one spot: one click, a double, or a
// triple.
type clickWatch struct {
	last  time.Time
	count int
	x, y  float64
}

// step records one press and returns the click count so far.
func (w *clickWatch) step(now time.Time, x, y float64) int {
	rapid := w.count > 0 && now.Sub(w.last) <= clickGap &&
		math.Abs(x-w.x) <= clickSlop && math.Abs(y-w.y) <= clickSlop

	if rapid {
		w.count++
	} else {
		w.count = 1
	}

	w.last, w.x, w.y = now, x, y

	return w.count
}
