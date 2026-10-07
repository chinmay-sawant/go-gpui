//go:build !linux && !windows

package collector

import (
	"runtime"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// liveSource reports that this platform has no collector. The UI labels the
// sections unavailable instead of showing zeroes.
func liveSource() domain.Source {
	return unsupported{reason: runtime.GOOS + " has no live collector in this build"}
}
