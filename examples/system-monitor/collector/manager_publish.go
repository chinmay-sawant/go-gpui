package collector

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// publishSample merges one raw sample with the previous one and stores it as
// the latest reading. The manager stamps the sample, so every source shares
// one monotonic timebase. Results from an old generation are dropped.
func (m *Manager) publishSample(gen uint64, name string, raw domain.Sample, stamp domain.Stamp) {
	raw.Stamp = stamp

	m.mu.Lock()
	if gen != m.gen {
		m.mu.Unlock()

		return
	}

	s := raw
	if m.haveRead {
		s = domain.Compute(m.prev, raw)
	}

	// A source without IO counters reports zero counters, which would turn
	// into a misleading 0 B/s rate. Keep those values unavailable instead.
	if !m.cap.DiskIO {
		for i := range s.Disks {
			s.Disks[i].ReadRate = domain.Value{}
			s.Disks[i].WriteRate = domain.Value{}
		}
	}

	m.prev = raw
	m.haveRead = true
	m.reading = s
	m.samples++
	m.lastSample = raw.Stamp.At
	m.runtime = m.readRuntimeLocked(raw.Stamp)
	m.pushRingsLocked(s)
	delete(m.errs, "sample")
	sink := m.sink
	m.mu.Unlock()

	if sink != nil {
		sink.RecordSample(s)
	}
}

// publishProcesses stores one process table read and turns cumulative CPU
// times into percent values against the previous read.
func (m *Manager) publishProcesses(gen uint64, name string, rows []domain.Process) {
	stamp := m.stampNow()

	m.mu.Lock()
	if gen != m.gen {
		m.mu.Unlock()

		return
	}

	if m.haveProcs {
		if elapsed, ok := domain.Elapsed(m.procStamp, stamp); ok {
			domain.ComputeProcesses(m.procPrev, rows, elapsed)
		}
	}

	prev := make(map[domain.ProcessIdentity]uint64, len(rows))
	for i := range rows {
		prev[rows[i].ID] = rows[i].CPUTime
	}

	m.procs = rows
	m.procStamp = stamp
	m.procPrev = prev
	m.haveProcs = true
	m.procSrc = name
	m.procSamples++
	m.lastProcess = stamp.At
	delete(m.errs, "processes")
	m.mu.Unlock()
}
