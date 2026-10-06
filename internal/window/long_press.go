package window

import (
	"math"
	"time"
)

const (
	longPressDelay = 450 * time.Millisecond
	longPressSlop  = 8.0
)

// longPressWatch tracks one held press: it arms on the down edge, fires once
// when the delay passes without a move past the slop, and remembers whether
// the screen claimed the press.
type longPressWatch struct {
	start   time.Time
	x, y    float64
	armed   bool
	fired   bool
	claimed bool
}

// arm starts the hold at a press point.
func (w *longPressWatch) arm(now time.Time, x, y float64) {
	w.start = now
	w.x, w.y = x, y
	w.armed, w.fired, w.claimed = true, false, false
}

// move cancels an unclaimed hold that drifts past the slop. A claimed hold
// keeps the gesture, so the move can extend the selection.
func (w *longPressWatch) move(x, y float64) {
	if !w.armed || w.claimed {
		return
	}

	if math.Abs(x-w.x) > longPressSlop || math.Abs(y-w.y) > longPressSlop {
		w.armed = false
	}
}

// cancel drops a pending hold, such as when a second finger joins.
func (w *longPressWatch) cancel() {
	w.armed = false
	w.claimed = false
}

// release ends the press. claimed stays set, so the lift that ends a claimed
// hold is not delivered as a tap.
func (w *longPressWatch) release() {
	w.armed = false
}

// due reports the first frame past the delay. fired spends the hold, so
// LongPress runs once per press.
func (w *longPressWatch) due(now time.Time) bool {
	return w.armed && !w.fired && now.Sub(w.start) >= longPressDelay
}
