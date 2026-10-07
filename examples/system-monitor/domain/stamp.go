package domain

import "time"

// Stamp is when a sample was taken. At is the wall clock time used for
// display and storage. Mono is the collector's monotonic timebase. Rates
// divide by Mono, so a wall clock step or a suspend cannot turn into a fake
// load spike.
type Stamp struct {
	At   time.Time
	Mono time.Duration
}

// Elapsed returns the monotonic distance between two stamps and whether a
// rate may divide by it. Same-instant and out-of-order stamps report false.
func Elapsed(prev, next Stamp) (time.Duration, bool) {
	d := next.Mono - prev.Mono
	if d <= 0 {
		return 0, false
	}

	return d, true
}

// DefaultGap is how far the wall clock may drift from the monotonic clock
// before Gap calls two stamps separated by a suspend or a clock step.
const DefaultGap = 5 * time.Second

// Gap reports whether the wall clock moved against the monotonic clock by at
// least threshold. A suspend and a wall clock step both look like this, and
// either way the two stamps belong to different rate eras. At values must
// carry a monotonic reading for this to work.
func Gap(prev, next Stamp, threshold time.Duration) bool {
	if threshold <= 0 {
		threshold = DefaultGap
	}

	wall := next.At.Round(0).Sub(prev.At.Round(0))
	drift := wall - (next.Mono - prev.Mono)

	return drift >= threshold || drift <= -threshold
}
