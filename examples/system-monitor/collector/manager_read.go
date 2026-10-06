package collector

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Reading returns the latest merged sample. Its slices are shared with the
// collector and must not be modified. Check HaveReading first.
func (m *Manager) Reading() domain.Sample {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.reading
}

// HaveReading reports whether a sample has been published. It is false until
// the first sample, and it is false again after a mode switch until the new
// source publishes.
func (m *Manager) HaveReading() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.haveRead
}

// Capabilities reports what the active source can read.
func (m *Manager) Capabilities() domain.Capabilities {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.cap
}

// SourceName returns the active source's name.
func (m *Manager) SourceName() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.source.Name()
}

// ProcessSnapshot is one published process table. Rows are shared with the
// collector and must not be modified; the UI sorts its own copy.
type ProcessSnapshot struct {
	Stamp  domain.Stamp
	Source string
	Rows   []domain.Process
}

// Processes returns the latest process table, or false when none has been
// published in the current mode.
func (m *Manager) Processes() (ProcessSnapshot, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return ProcessSnapshot{Stamp: m.procStamp, Source: m.procSrc, Rows: m.procs}, m.haveProcs
}
