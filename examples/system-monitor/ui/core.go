package ui

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
)

// Collector adapts a collector.Manager to the Source the UI polls. The
// manager keeps the samples; every call here reads its cached state and
// never waits on a worker.
type Collector struct {
	mgr *collector.Manager
}

// NewCollector wraps a started manager.
func NewCollector(mgr *collector.Manager) *Collector {
	return &Collector{mgr: mgr}
}

// Summary converts the latest merged sample. It reports false before the
// first sample of the current mode.
func (c *Collector) Summary(context.Context) (Summary, bool) {
	if !c.mgr.HaveReading() {
		return Summary{}, false
	}

	return toSummary(c.mgr.Reading()), true
}

// Processes converts the latest process table.
func (c *Collector) Processes(context.Context) (ProcSnapshot, bool) {
	snap, ok := c.mgr.Processes()
	if !ok {
		return ProcSnapshot{}, false
	}

	out := ProcSnapshot{
		At:    snap.Stamp.At,
		Procs: make([]Process, 0, len(snap.Rows)),
	}

	for _, p := range snap.Rows {
		out.Procs = append(out.Procs, toProcess(p))
	}

	return out, true
}

// Track parses an identity and asks the manager to follow it. The manager
// drops results for an older selection, so a stale detail never attaches.
func (c *Collector) Track(id string) {
	c.mgr.Track(parseID(id))
}

// Tracked converts the tracked process state.
func (c *Collector) Tracked(context.Context) Tracked {
	return toTracked(c.mgr.Tracked())
}
