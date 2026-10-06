package storage

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// fixtureSamples builds three labeled fixture samples.
func fixtureSamples() []domain.Sample {
	base := time.Unix(1_700_000_000, 0)

	out := make([]domain.Sample, 0, 3)

	for i := range 3 {
		s := domain.Sample{
			Stamp:      domain.Stamp{At: base.Add(time.Duration(i) * time.Second), Mono: time.Duration(i) * time.Second},
			CPUPercent: domain.Percent(float64(i + 1)),
		}
		s.CPUTotal = domain.CPUTimes{Busy: uint64(i + 1), Total: 10}
		s.Mem.Used = domain.Bytes(100)
		out = append(out, s)
	}

	return out
}

// fixture returns a labeled fixture value.
func fixture(version int) domain.Fixture {
	base := time.Unix(1_700_000_000, 0)

	return domain.Fixture{
		Version: version,
		Name:    "dummy history",
		Source:  "dummy",
		Note:    "seeded fixture",
		Started: base,
		Step:    time.Second,
		Samples: fixtureSamples(),
	}
}
