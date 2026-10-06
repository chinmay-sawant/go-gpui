package game

import "time"

// Clock converts wall time into bounded fixed steps. The UI calls Advance
// once per window frame, then runs Step for each step it returns.
type Clock struct {
	last time.Time
	acc  time.Duration
}

// StallLimit is the longest frame gap that still simulates. A longer gap
// pauses play instead of fast-forwarding after focus loss or system sleep.
const StallLimit = 500 * time.Millisecond

// MaxCatchUp bounds the fixed steps one Advance can return.
const MaxCatchUp = 5

// Advance returns how many fixed steps to run now, at most MaxCatchUp.
// The first call only starts the clock. A gap at or over StallLimit
// discards the elapsed time and returns 0.
func (c *Clock) Advance(now time.Time) int { return 0 }

// Reset discards elapsed wall time, so the next Advance runs no steps.
func (c *Clock) Reset(now time.Time) {}
