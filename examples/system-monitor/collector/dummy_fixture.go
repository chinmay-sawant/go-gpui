package collector

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// SeedVersion is the fixture version written to storage. A bump replaces the
// stored dummy session once.
const SeedVersion = 1

// DummyFixture builds the labeled recording a fresh database is seeded from:
// 120 samples ending at now. The fixture source is marked "dummy" in storage,
// so it is never mistaken for live history.
func DummyFixture(seed int64, now time.Time, count int, step time.Duration) domain.Fixture {
	d := NewDummy(DummyOptions{Seed: seed, Processes: DefaultDummyProcesses})
	samples := d.Historical(now, count, step)

	started := now
	if len(samples) > 0 {
		started = samples[0].Stamp.At
	}

	return domain.Fixture{
		Version: SeedVersion,
		Name:    "dummy history",
		Source:  "dummy",
		Note:    "seeded fixture, not a live recording",
		Started: started,
		Step:    step,
		Samples: samples,
	}
}
