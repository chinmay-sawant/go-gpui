package collector

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Tracked is the state of the tracked process detail view.
type Tracked struct {
	ID      domain.ProcessIdentity
	Detail  domain.ProcessDetail
	Loading bool
	Err     string
	Stamp   domain.Stamp
}

// Tracked returns the state of the selected process. Loading is true between
// Track and the next detail result.
func (m *Manager) Tracked() Tracked {
	m.mu.Lock()
	defer m.mu.Unlock()

	return Tracked{
		ID:      m.tracked,
		Detail:  m.detail,
		Loading: m.detailLoad,
		Err:     m.detailErr,
		Stamp:   m.detailStamp,
	}
}

// Track selects one process for the detail view. The zero identity clears the
// selection. A result for the previous selection is dropped.
func (m *Manager) Track(id domain.ProcessIdentity) {
	m.mu.Lock()
	m.tracked = id
	m.haveDetail = false
	m.detailPrev = domain.ProcessDetail{}
	m.detailLoad = id != (domain.ProcessIdentity{})
	m.detailErr = ""
	m.mu.Unlock()
}
