package collector

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// ringLocked returns the ring for key, creating it and marking it fresh.
func (m *Manager) ringLocked(key domain.SeriesKey) *domain.Ring {
	if m.seen == nil {
		m.seen = make(map[domain.SeriesKey]time.Time)
	}

	r, ok := m.rings[key]
	if !ok {
		r = domain.NewRing(m.opts.Capacity)
		m.rings[key] = r
	}
	m.seen[key] = time.Now()

	return r
}

// pruneRingsLocked drops buffers the source stopped reporting. A prefill push
// marks every key it touches, so history survives.
func (m *Manager) pruneRingsLocked(now time.Time) {
	for key, seen := range m.seen {
		if now.Sub(seen) > staleRing {
			delete(m.rings, key)
			delete(m.seen, key)
		}
	}
}
