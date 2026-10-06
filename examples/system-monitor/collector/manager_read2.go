package collector

import (
	"sort"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Runtime returns this process's own metrics, sampled with each summary.
func (m *Manager) Runtime() domain.Runtime {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.runtime
}

// Errors returns the latest failure of each step, in a stable order. A step
// that succeeds again is removed, so an empty slice means every step is
// healthy.
func (m *Manager) Errors() []Error {
	m.mu.Lock()
	defer m.mu.Unlock()

	out := make([]Error, 0, len(m.errs))
	for _, e := range m.errs {
		out = append(out, e)
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Step < out[j].Step })

	return out
}

// SetSink installs or removes the recording sink. A nil sink stops recording.
// The sink is called outside the manager lock and must not block.
func (m *Manager) SetSink(s Sink) {
	m.mu.Lock()
	m.sink = s
	m.mu.Unlock()
}

// Stats returns manager and worker pool counters for diagnostics.
func (m *Manager) Stats() Stats {
	m.mu.Lock()
	defer m.mu.Unlock()

	return Stats{
		Mode:           m.mode,
		Source:         m.source.Name(),
		Generation:     m.gen,
		Samples:        m.samples,
		SampleErrors:   m.sampleErrs,
		ProcessSamples: m.procSamples,
		ProcessErrors:  m.procErrs,
		DetailSamples:  m.detailCount,
		DetailErrors:   m.detailErrs,
		Skipped:        m.skipped,
		LastSample:     m.lastSample,
		LastProcess:    m.lastProcess,
		LastDetail:     m.lastDetailAt,
		Pool:           m.pool.Stats(),
	}
}

// Stats is a point-in-time view of the collector.
type Stats struct {
	Mode           Mode
	Source         string
	Generation     uint64
	Samples        uint64
	SampleErrors   uint64
	ProcessSamples uint64
	ProcessErrors  uint64
	DetailSamples  uint64
	DetailErrors   uint64
	Skipped        uint64
	LastSample     time.Time
	LastProcess    time.Time
	LastDetail     time.Time
	Pool           PoolStats
}
