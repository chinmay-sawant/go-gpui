package collector

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// DummyOptions configures the reproducible dummy source. The same seed and
// the same call sequence always produce the same samples.
type DummyOptions struct {
	Seed      int64
	Processes int
	// Stamp returns the current stamp. Nil uses wall and process time.
	Stamp func() domain.Stamp
}

// Dummy is the labeled fixture source. It reads no host state, so the example
// runs without permissions, and its history never mixes with live readings.
type Dummy struct {
	opts DummyOptions
}

// NewDummy returns a fixture source.
func NewDummy(opts DummyOptions) *Dummy {
	if opts.Processes <= 0 {
		opts.Processes = DefaultDummyProcesses
	}
	if opts.Stamp == nil {
		start := time.Now()
		opts.Stamp = func() domain.Stamp {
			return domain.Stamp{At: time.Now(), Mono: time.Since(start)}
		}
	}

	return &Dummy{opts: opts}
}

// NewStress returns the 10,000 process fixture used by stress tests and the
// -stress flag.
func NewStress(seed int64) *Dummy {
	return NewDummy(DummyOptions{Seed: seed, Processes: StressProcesses})
}

// Name labels the source.
func (d *Dummy) Name() string { return "dummy" }

// Capabilities reports the fixture surface.
func (d *Dummy) Capabilities() domain.Capabilities { return domain.Capabilities{} }

// Sample reads the next fixture sample.
func (d *Dummy) Sample(ctx context.Context) (domain.Sample, error) {
	return domain.Sample{}, nil
}

// Processes reads the fixture process table.
func (d *Dummy) Processes(ctx context.Context) ([]domain.Process, error) {
	return nil, nil
}

// Detail reads one fixture process.
func (d *Dummy) Detail(ctx context.Context, id domain.ProcessIdentity) (domain.ProcessDetail, error) {
	return domain.ProcessDetail{}, nil
}

// Historical returns count samples ending at start, spaced step apart.
func (d *Dummy) Historical(start time.Time, count int, step time.Duration) []domain.Sample {
	return nil
}
