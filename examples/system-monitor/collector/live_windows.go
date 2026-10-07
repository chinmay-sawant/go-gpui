//go:build windows

package collector

import (
	"os"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// liveSource returns the kernel32 source on Windows.
func liveSource() domain.Source { return newWindowsSource() }

// windowsSource reads system counters through kernel32 and psapi. Windows
// has no /proc equivalent that is this cheap, so the source reports only what
// those APIs expose and marks the rest unavailable.
type windowsSource struct {
	host string
}

func newWindowsSource() *windowsSource {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	return &windowsSource{host: host}
}

// Name labels the source.
func (s *windowsSource) Name() string { return "win32" }

// Capabilities reports what this adapter reads. The notes become the
// unavailable labels in the UI.
func (s *windowsSource) Capabilities() domain.Capabilities {
	return domain.Capabilities{
		PerCoreCPU:    false,
		Processes:     true,
		ProcessDetail: true,
		Disk:          true,
		DiskIO:        false,
		Net:           false,
		Load:          false,
		Sensors:       false,
		Handles:       false,
		Swap:          true,
		Notes: []string{
			"per-core CPU needs NtQuerySystemInformation and is unavailable",
			"disk IO rates need PDH and are unavailable",
			"network counters need Iphlpapi and are unavailable",
			"temperatures need WMI and are unavailable",
			"command lines need WMI; the executable path is shown instead",
			"protected processes report no CPU time or start time",
		},
	}
}
