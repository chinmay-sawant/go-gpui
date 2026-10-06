package collector

import (
	"sort"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Series returns a copy of one graph buffer's points, oldest first.
func (m *Manager) Series(key domain.SeriesKey) (domain.Series, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	r, ok := m.rings[key]
	if !ok {
		return domain.Series{}, false
	}

	return domain.Series{Key: key, Points: r.Points()}, true
}

// Keys returns every graph buffer key, sorted by metric and device.
func (m *Manager) Keys() []domain.SeriesKey {
	m.mu.Lock()
	defer m.mu.Unlock()

	keys := make([]domain.SeriesKey, 0, len(m.rings))
	for key := range m.rings {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Metric != keys[j].Metric {
			return keys[i].Metric < keys[j].Metric
		}

		return keys[i].Device < keys[j].Device
	})

	return keys
}
