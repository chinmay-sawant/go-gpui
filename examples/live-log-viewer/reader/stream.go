package reader

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// StreamOptions configures a generated source. Seed, Key, and Start sequence
// make the output reproducible; Rate and Burst bound how much one Read may
// produce. A stream is not replayable, so every record skipped by an
// overload is reported through Batch.Lost.
type StreamOptions struct {
	Seed      int64
	Key       string
	Start     int64
	Count     int64
	Rate      int
	Burst     int
	MaxRecord int
	Policy    entry.Policy
	Clock     func() time.Time
	Sleep     func(context.Context, time.Duration) error
}

// Stream is the dummy and burst generator.
type Stream struct {
	opts   StreamOptions
	pol    entry.Policy
	seq    int64
	last   time.Time
	carry  int64
	lost   int64
	done   bool
	closed bool
}

// NewStream prepares a generated source with the policy's batch limits.
func NewStream(o StreamOptions) *Stream {
	if o.Policy == (entry.Policy{}) {
		o.Policy = entry.DefaultPolicy()
	}

	o.Policy = o.Policy.ForStream()

	if o.Burst < 1 {
		o.Burst = o.Policy.BatchRecords
	}

	if o.MaxRecord < 1 {
		o.MaxRecord = o.Policy.MaxRecord
	}

	if o.Start < 1 {
		o.Start = 1
	}

	return &Stream{opts: o, pol: o.Policy, seq: o.Start}
}

// NewDummy returns the default demo generator.
func NewDummy(o StreamOptions) *Stream {
	if o.Key == "" {
		o.Key = "dummy"
	}

	return NewStream(o)
}

// NewBurst returns a bounded stream that runs Count records and reports Done.
func NewBurst(o StreamOptions) *Stream {
	if o.Key == "" {
		o.Key = "burst"
	}

	return NewStream(o)
}
