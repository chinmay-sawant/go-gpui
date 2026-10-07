package reader

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Read produces the next bounded batch. A live rate that outran the consumer
// is reported as loss, never dropped silently.
func (s *Stream) Read(ctx context.Context) (Batch, error) {
	if err := ctx.Err(); err != nil {
		return Batch{}, err
	}

	if s.closed {
		return Batch{}, ErrClosed
	}

	if s.done {
		return Batch{Generation: 1, Position: s.seq, State: entry.StateDone, Done: true}, nil
	}

	n := s.due()
	if n == 0 {
		if err := s.wait(ctx, s.until()); err != nil {
			return Batch{}, err
		}

		return Batch{Generation: 1, Position: s.seq, State: entry.StateLive, Lost: s.takeLost()}, nil
	}

	if s.opts.Count > 0 && s.seq+int64(n) > s.opts.Count+1 {
		n = int(s.opts.Count + 1 - s.seq)
	}

	if n < 0 {
		n = 0
	}

	recs := DummyRecords(s.opts.Seed, s.opts.Key, s.seq, n, s.opts.MaxRecord)
	s.seq += int64(n)

	b := Batch{
		Records: recs, Generation: 1, Position: s.seq,
		State: entry.StateLive, More: n >= s.opts.Burst, Lost: s.takeLost(),
	}

	if s.opts.Count > 0 && s.seq > s.opts.Count {
		s.done = true
		b.Done, b.State = true, entry.StateDone
	}

	return b, nil
}
