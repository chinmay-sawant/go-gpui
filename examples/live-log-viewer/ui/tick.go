package ui

import (
	"context"
	"time"
)

// Tick runs once per frame before the window draws. It drains worker
// results inside a bounded budget, dispatches a due search, retries a
// coalesced snapshot, and refreshes the sidebar on a slow clock. It never
// waits on the feed; it redraws only when something changed.
func (a *App) Tick(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	start := time.Now()
	now := start
	changed := a.drain(now)
	a.retry()

	if _, _, ok := a.filters.Ready(now); ok {
		a.loadPage(intentFilter)
	}

	if now.Sub(a.sourcesDue) >= sourcesInterval {
		a.sourcesDue = now
		a.refreshSources()
	}

	if a.perf {
		a.lastTick = time.Since(start)
	}

	if changed {
		return a.draw(ctx)
	}

	return nil
}

// drain applies at most DrainBudget results or DrainTime of wall time.
func (a *App) drain(now time.Time) bool {
	changed := false
	deadline := time.Now().Add(DrainTime)

	for i := 0; i < DrainBudget; i++ {
		select {
		case o := <-a.outs:
			if a.apply(o, now) {
				changed = true
			}
		default:
			return changed
		}

		if time.Now().After(deadline) {
			break
		}
	}

	return changed
}
