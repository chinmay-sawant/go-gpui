package collector

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Historical builds count samples ending at start, spaced step apart, from
// this source's simulation. The counters continue where the history ends, so
// the first live sample after a prefill has a real previous reading. Each
// sample after the first carries rates computed the same way the manager
// computes them.
func (d *Dummy) Historical(start time.Time, count int, step time.Duration) []domain.Sample {
	if count <= 0 {
		return nil
	}
	if step <= 0 {
		step = DefaultHistoryStep
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	out := make([]domain.Sample, 0, count)
	var prev domain.Sample

	for i := range count {
		back := time.Duration(count-1-i) * step
		stamp := domain.Stamp{At: start.Add(-back), Mono: -back}
		raw := d.advance(stamp)

		s := raw
		if i > 0 {
			s = domain.Compute(prev, raw)
		}
		out = append(out, s)
		prev = raw
	}

	return out
}
