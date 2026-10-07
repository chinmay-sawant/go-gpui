//go:build linux

package collector

import (
	"os"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// liveSource returns the /proc source on Linux.
func liveSource() domain.Source { return newProcSource() }

// procSource reads machine state from /proc and /sys. Every file read is
// best effort: a process that exits mid-listing is skipped, and an optional
// file that is missing leaves its values unavailable instead of zero.
type procSource struct {
	host         string
	pageSize     int
	hasSensors   bool
	hasDiskIO    bool
	hasDiskSpace bool
	hasNet       bool

	users map[int32]string
}

func newProcSource() *procSource {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	return &procSource{
		host:         host,
		pageSize:     os.Getpagesize(),
		hasSensors:   pathExists("/sys/class/hwmon"),
		hasDiskIO:    pathExists("/proc/diskstats"),
		hasDiskSpace: pathExists("/proc/mounts"),
		hasNet:       pathExists("/proc/net/dev"),
		users:        make(map[int32]string),
	}
}

// Name labels the source.
func (s *procSource) Name() string { return "procfs" }
