package collector

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// buildSource returns the source for a mode. Callers hold m.mu.
func (m *Manager) buildSource(mode Mode) domain.Source {
	if m.opts.Source != nil {
		return m.opts.Source
	}
	if mode == ModeLive {
		if m.live == nil {
			m.live = NewLive()
		}

		return m.live
	}

	return NewDummy(DummyOptions{
		Seed:      m.opts.Seed,
		Processes: m.opts.processCount(),
		Stamp:     m.stampNow,
	})
}

// stampNow returns the manager's current stamp. Mono starts at zero in Start.
func (m *Manager) stampNow() domain.Stamp {
	at := time.Now()

	if m.start.IsZero() {
		return domain.Stamp{At: at}
	}

	return domain.Stamp{At: at, Mono: time.Since(m.start)}
}
