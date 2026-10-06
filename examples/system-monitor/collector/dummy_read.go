package collector

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Capabilities reports the fixture surface. Handle counts are the one thing
// it does not model, so the UI shows them unavailable.
func (d *Dummy) Capabilities() domain.Capabilities {
	return domain.Capabilities{
		PerCoreCPU:    true,
		Processes:     true,
		ProcessDetail: true,
		Disk:          true,
		DiskIO:        true,
		Net:           true,
		Load:          true,
		Sensors:       true,
		Handles:       false,
		Swap:          true,
		Notes:         []string{"dummy fixture, no host access"},
	}
}

// Sample reads the next fixture sample.
func (d *Dummy) Sample(ctx context.Context) (domain.Sample, error) {
	if err := ctx.Err(); err != nil {
		return domain.Sample{}, err
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	return d.advance(d.opts.Stamp()), nil
}
