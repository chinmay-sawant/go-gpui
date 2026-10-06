package domain

import (
	"context"
	"errors"
)

// ErrUnsupported reports that a source cannot read something on this platform.
// Callers show an unavailable value, never a zero.
var ErrUnsupported = errors.New("system-monitor: unsupported on this platform")

// Capabilities says what a source can read. A false field means every related
// value is invalid, so the UI labels the section instead of drawing an empty
// graph. Notes carries short, user-visible limitations such as "process CPU
// needs WMI on this platform".
type Capabilities struct {
	PerCoreCPU    bool
	Processes     bool
	ProcessDetail bool
	Disk          bool
	DiskIO        bool
	Net           bool
	Load          bool
	Sensors       bool
	Handles       bool
	Swap          bool
	Notes         []string
}

// Source reads the machine. Methods run on the collector's worker goroutines
// and may be called concurrently, at most one call of each kind at a time. A
// source returns when ctx is done; it does not keep the context, start
// long-lived goroutines, or panic. Sample fills raw fields only, because the
// manager computes rates. A source that a platform cannot support returns
// ErrUnsupported.
type Source interface {
	// Name labels the source in the UI, for example "dummy" or "procfs".
	Name() string
	// Capabilities reports what the source can read.
	Capabilities() Capabilities
	// Sample reads machine-level counters and readings.
	Sample(ctx context.Context) (Sample, error)
	// Processes reads the process table.
	Processes(ctx context.Context) ([]Process, error)
	// Detail reads one process in full.
	Detail(ctx context.Context, id ProcessIdentity) (ProcessDetail, error)
}
