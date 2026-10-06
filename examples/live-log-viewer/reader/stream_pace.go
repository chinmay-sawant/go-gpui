package reader

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Flush has nothing to emit: generated records end at a record boundary.
func (s *Stream) Flush() []entry.RawRecord { return nil }

// Replayable reports false: a generated record cannot be reread.
func (s *Stream) Replayable() bool { return false }

// Close stops the generator.
func (s *Stream) Close() error {
	s.closed = true

	return nil
}

// due returns how many whole records are due now and counts the ones the
// consumer cannot take as lost.
func (s *Stream) due() int {
	if s.opts.Rate <= 0 {
		return s.opts.Burst
	}

	now := s.clock()
	if s.last.IsZero() {
		s.last = now
	}

	if elapsed := now.Sub(s.last); elapsed > 0 {
		s.carry += elapsed.Nanoseconds() * int64(s.opts.Rate)
		s.last = now
	}

	due := s.carry / int64(time.Second)
	s.carry -= due * int64(time.Second)

	if max := int64(s.opts.Burst); due > max {
		s.lost += due - max
		due = max
	}

	return int(due)
}

// until is the wait before the next record is due.
func (s *Stream) until() time.Duration {
	if s.opts.Rate <= 0 {
		return s.pol.Poll
	}

	d := time.Duration((int64(time.Second) - s.carry) / int64(s.opts.Rate))
	if d < time.Millisecond {
		d = time.Millisecond
	}

	return d
}

func (s *Stream) clock() time.Time {
	if s.opts.Clock != nil {
		return s.opts.Clock()
	}

	return time.Now()
}

func (s *Stream) wait(ctx context.Context, d time.Duration) error {
	if s.opts.Sleep != nil {
		return s.opts.Sleep(ctx, d)
	}

	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func (s *Stream) takeLost() int64 {
	lost := s.lost
	s.lost = 0

	return lost
}
