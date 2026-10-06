package domain

import "time"

// MaxSampleGap is the longest elapsed time a rate sample may span. A
// longer gap means the transfer was suspended, so the sample is dropped.
const MaxSampleGap = 30 * time.Second

// Rate returns the bytes per second between two progress samples. It
// returns ok=false for the first sample, for a clock that does not advance,
// for negative progress, and for a gap longer than maxGap. Dropping the
// suspended gap keeps a resume from poisoning the average.
func Rate(prevDone int64, prevAt time.Time, done int64, at time.Time, maxGap time.Duration) (float64, bool) {
	if prevAt.IsZero() || at.IsZero() || !at.After(prevAt) {
		return 0, false
	}

	gap := at.Sub(prevAt)
	if maxGap > 0 && gap > maxGap {
		return 0, false
	}

	delta := done - prevDone
	if delta < 0 {
		return 0, false
	}

	return float64(delta) / gap.Seconds(), true
}
