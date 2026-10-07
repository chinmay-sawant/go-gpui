package scene

import "time"

// Options configure New.
type Options struct {
	// Dark starts on the dark theme; a saved preference replaces it when
	// the store answers.
	Dark bool
	// Status is the first status line, for a startup notice.
	Status string
	// Perf records pipeline timing, dirty counts, and allocs in the page
	// stats, and lets the window sample frame times.
	Perf bool
	// Now overrides the clock, for tests and replay-style runs.
	Now func() time.Time
	// Focused reports whether the window has focus; false pauses play.
	Focused func() bool
	// Stepper converts wall time into fixed steps; nil uses game.Clock.
	Stepper Stepper
}
