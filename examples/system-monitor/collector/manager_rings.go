package collector

import (
	"strconv"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// staleRing is how long a series keeps its points after the source stops
// reporting it, so an unplugged interface disappears from the UI.
const staleRing = 5 * time.Minute

// pushRingsLocked appends one sample to every graph buffer and drops buffers
// the source no longer reports.
func (m *Manager) pushRingsLocked(s domain.Sample) {
	gap := s.Gap

	point := func(v domain.Value) domain.Point {
		return domain.Point{At: s.Stamp.At, Value: v.N, Valid: v.Valid, Gap: gap}
	}

	now := time.Now()

	m.ringLocked(domain.SeriesKey{Metric: domain.MetricCPU}).Push(point(s.CPUPercent))
	m.ringLocked(domain.SeriesKey{Metric: domain.MetricMemPercent}).Push(point(s.Mem.UsedPercent))
	m.ringLocked(domain.SeriesKey{Metric: domain.MetricProcs}).Push(point(s.Procs))
	m.ringLocked(domain.SeriesKey{Metric: domain.MetricThreads}).Push(point(s.Threads))

	for i, core := range s.CorePercent {
		key := domain.SeriesKey{Metric: domain.MetricCPUCore, Device: strconv.Itoa(i)}
		m.ringLocked(key).Push(point(core))
	}
	for _, d := range s.Disks {
		m.ringLocked(domain.SeriesKey{Metric: domain.MetricDiskRead, Device: d.Device}).Push(point(d.ReadRate))
		m.ringLocked(domain.SeriesKey{Metric: domain.MetricDiskWrite, Device: d.Device}).Push(point(d.WriteRate))
	}
	for _, n := range s.Nets {
		m.ringLocked(domain.SeriesKey{Metric: domain.MetricNetRX, Device: n.Name}).Push(point(n.RXRate))
		m.ringLocked(domain.SeriesKey{Metric: domain.MetricNetTX, Device: n.Name}).Push(point(n.TXRate))
	}

	m.pruneRingsLocked(now)
}
