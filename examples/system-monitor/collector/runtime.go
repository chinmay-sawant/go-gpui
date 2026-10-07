package collector

import (
	"runtime"
	"runtime/metrics"
	"sync"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// readRuntimeLocked samples this process's metrics. Callers hold m.mu.
func (m *Manager) readRuntimeLocked(stamp domain.Stamp) domain.Runtime {
	r := domain.Runtime{Stamp: stamp}

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	r.Goroutines = runtime.NumGoroutine()
	r.HeapBytes = ms.HeapAlloc
	r.HeapGoal = ms.NextGC
	r.SysBytes = ms.Sys
	r.AllocBytes = ms.TotalAlloc
	r.GCCount = ms.NumGC
	r.GCPause = time.Duration(ms.PauseTotalNs)
	if !m.start.IsZero() {
		r.Uptime = time.Since(m.start)
	}

	if cpu, ok := cpuSeconds(); ok {
		if m.rt.ok {
			prev := domain.Stamp{At: stamp.At, Mono: m.rt.mono}
			if elapsed, ok := domain.Elapsed(prev, stamp); ok && cpu >= m.rt.cpu {
				r.CPUPercent = domain.Percent((cpu - m.rt.cpu) / elapsed.Seconds() * 100)
			}
		}

		m.rt = rtSample{ok: true, cpu: cpu, mono: stamp.Mono}
	}

	return r
}

// cpuMetricName finds the cumulative process CPU metric once.
var cpuMetricName = sync.OnceValue(func() string {
	for _, d := range metrics.All() {
		if d.Name == "/cpu/classes/total:cpu-seconds" {
			return d.Name
		}
	}

	return ""
})

// cpuSeconds returns this process's cumulative CPU time when the runtime
// exposes it on this platform.
func cpuSeconds() (float64, bool) {
	name := cpuMetricName()
	if name == "" {
		return 0, false
	}

	samples := []metrics.Sample{{Name: name}}
	metrics.Read(samples)
	if samples[0].Value.Kind() == metrics.KindBad {
		return 0, false
	}

	return samples[0].Value.Float64(), true
}
