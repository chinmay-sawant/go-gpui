package collector

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// NewLive returns the platform's live source. Platforms without an adapter
// get a source that reports everything as unsupported, so the UI shows
// unavailable values instead of zeroes.
func NewLive() domain.Source { return unsupported{} }

// unsupported stands in for a platform with no live collector.
type unsupported struct{}

func (unsupported) Name() string { return "unsupported" }

func (unsupported) Capabilities() domain.Capabilities { return domain.Capabilities{} }

func (unsupported) Sample(ctx context.Context) (domain.Sample, error) {
	return domain.Sample{}, domain.ErrUnsupported
}

func (unsupported) Processes(ctx context.Context) ([]domain.Process, error) {
	return nil, domain.ErrUnsupported
}

func (unsupported) Detail(ctx context.Context, id domain.ProcessIdentity) (domain.ProcessDetail, error) {
	return domain.ProcessDetail{}, domain.ErrUnsupported
}

// SeedVersion is the fixture version written to storage. A bump replaces the
// stored dummy session once.
const SeedVersion = 1

// DummyFixture builds the labeled 120 sample recording a fresh database is
// seeded from.
func DummyFixture(seed int64, now time.Time, count int, step time.Duration) domain.Fixture {
	return domain.Fixture{Version: SeedVersion, Name: "dummy history", Source: "dummy", Started: now}
}
