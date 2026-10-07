// Package collector samples the machine and publishes immutable snapshots for
// the UI. It owns the worker pool, the sampling cadence, rate computation, and
// the graph buffers. The UI never reads the host itself: it reads the latest
// published values and asks for a process detail when the user selects one.
//
// A manager runs two loops, one for cheap summary counters about once a
// second and one for the expensive process table less often. Both submit to a
// bounded pool, so a slow source skips a tick instead of piling up work. A
// tick never starts a goroutine per process.
package collector
