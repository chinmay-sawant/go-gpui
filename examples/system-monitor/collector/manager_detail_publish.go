package collector

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// publishDetail stores one tracked process read. A result for a different
// identity than the current selection is dropped.
func (m *Manager) publishDetail(gen uint64, name string, id domain.ProcessIdentity, d domain.ProcessDetail) {
	stamp := m.stampNow()

	m.mu.Lock()
	if gen != m.gen || id != m.tracked {
		m.mu.Unlock()

		return
	}

	if m.haveDetail && m.detailPrev.ID == d.ID {
		if elapsed, ok := domain.Elapsed(m.detailStamp, stamp); ok {
			d = domain.ComputeDetail(m.detailPrev, d, elapsed)
		}
	}

	m.detail = d
	m.detailPrev = d
	m.haveDetail = true
	m.detailLoad = false
	m.detailErr = ""
	m.detailSrc = name
	m.detailStamp = stamp
	m.detailCount++
	m.lastDetailAt = stamp.At
	m.mu.Unlock()
}
