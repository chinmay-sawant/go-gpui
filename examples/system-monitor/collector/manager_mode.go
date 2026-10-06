package collector

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// SetMode switches between dummy and live sampling. It invalidates results
// from the old mode, drops the rate baselines, and clears the graph buffers,
// so dummy history never mixes with live readings. In dummy mode it fills the
// buffers with the fixture history.
func (m *Manager) SetMode(mode Mode) error {
	m.mu.Lock()
	if m.closed {
		m.mu.Unlock()

		return ErrClosed
	}
	if mode == m.mode {
		m.mu.Unlock()

		return nil
	}

	m.mode = mode
	m.gen++
	m.source = m.buildSource(mode)
	m.cap = m.source.Capabilities()
	m.reading = domain.Sample{}
	m.haveRead = false
	m.prev = domain.Sample{}
	m.procs = nil
	m.procPrev = nil
	m.haveProcs = false
	m.procSrc = ""
	m.tracked = domain.ProcessIdentity{}
	m.detail = domain.ProcessDetail{}
	m.detailPrev = domain.ProcessDetail{}
	m.haveDetail = false
	m.detailLoad = false
	m.detailErr = ""
	m.rings = make(map[domain.SeriesKey]*domain.Ring)
	m.seen = nil
	m.errs = make(map[string]Error)
	m.rt = rtSample{}

	if mode == ModeDummy && m.started {
		m.prefillLocked()
	}
	m.mu.Unlock()

	return nil
}

// generation returns the current mode generation.
func (m *Manager) generation() uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.gen
}

// stillCurrent reports whether gen is still the live generation.
func (m *Manager) stillCurrent(gen uint64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return gen == m.gen
}

// currentSource returns the source and its name.
func (m *Manager) currentSource() (domain.Source, string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.source, m.source.Name()
}
