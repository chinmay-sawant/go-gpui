package game

import "time"

// Clock converts wall time into bounded fixed steps. The UI calls Advance
// once per window frame, then runs Step once per returned step.
type Clock struct {
	last time.Time
	acc  time.Duration
}

// StallLimit is the longest frame gap that still simulates. A longer gap
// pauses play instead of fast-forwarding after focus loss or sleep.
const StallLimit = 500 * time.Millisecond

// MaxCatchUp bounds the fixed steps one Advance can return.
const MaxCatchUp = 5

// Advance returns how many fixed steps to run now, at most MaxCatchUp.
// The first call only starts the clock. A gap at or over StallLimit
// discards the elapsed time and returns 0.
func (c *Clock) Advance(now time.Time) int {
	if c.last.IsZero() {
		c.last = now

		return 0
	}

	dt := now.Sub(c.last)
	c.last = now

	if dt <= 0 {
		return 0
	}

	if dt >= StallLimit {
		c.acc = 0

		return 0
	}

	c.acc += dt
	if budget := FixedStep * MaxCatchUp; c.acc > budget {
		c.acc = budget
	}

	n := 0
	for c.acc >= FixedStep && n < MaxCatchUp {
		c.acc -= FixedStep
		n++
	}

	return n
}

// Reset discards elapsed wall time, so the next Advance runs no steps.
// The UI calls it on focus loss and focus gain.
func (c *Clock) Reset(now time.Time) {
	c.last = now
	c.acc = 0
}
