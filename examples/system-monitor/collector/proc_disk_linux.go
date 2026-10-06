//go:build linux

package collector

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// diskIO is cumulative byte counts for one block device.
type diskIO struct {
	read  uint64
	write uint64
}

// readDisks joins /proc/diskstats counters with space from the mounts that
// live on a whole block device.
func (s *procSource) readDisks(ctx context.Context) []domain.Disk {
	if !s.hasDiskIO && !s.hasDiskSpace {
		return nil
	}

	io := parseDiskStats(readFile("/proc/diskstats"))
	out := make([]domain.Disk, 0, len(io))
	seen := make(map[string]bool)

	for _, m := range parseMounts(readFile("/proc/mounts")) {
		if err := ctx.Err(); err != nil {
			return out
		}
		if !strings.HasPrefix(m.device, "/dev/") {
			continue
		}

		dev := wholeDiskName(filepath.Base(m.device))
		if dev == "" || seen[dev] {
			continue
		}
		seen[dev] = true

		d := domain.Disk{Device: dev, Mount: m.mount}
		if counters, ok := io[dev]; ok {
			d.ReadBytes, d.WriteBytes = counters.read, counters.write
		}
		if total, free, err := statfs(m.mount); err == nil {
			d.Total = domain.Bytes(float64(total))
			d.Free = domain.Bytes(float64(free))
		}

		out = append(out, d)
	}

	return out
}
