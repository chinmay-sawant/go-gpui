package collector

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// NewLive returns the platform's live source. Platforms without an adapter
// get a source that reports everything as unsupported, so the UI shows
// unavailable values instead of zeroes.
func NewLive() domain.Source { return liveSource() }

// unsupported stands in for a platform with no live collector.
type unsupported struct{ reason string }

func (u unsupported) Name() string { return "unsupported" }

func (u unsupported) Capabilities() domain.Capabilities {
	note := u.reason
	if note == "" {
		note = "no live collector on this platform"
	}

	return domain.Capabilities{Notes: []string{note}}
}

func (unsupported) Sample(ctx context.Context) (domain.Sample, error) {
	return domain.Sample{}, domain.ErrUnsupported
}

func (unsupported) Processes(ctx context.Context) ([]domain.Process, error) {
	return nil, domain.ErrUnsupported
}

func (unsupported) Detail(ctx context.Context, id domain.ProcessIdentity) (domain.ProcessDetail, error) {
	return domain.ProcessDetail{}, domain.ErrUnsupported
}
