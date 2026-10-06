package ui

import (
	"context"
	"time"
)

// Tick is the Page.SetTick callback. It drains worker results within a
// bounded budget, relayouts only when geometry changed, and otherwise
// repaints the retained progress operations in place.
func (a *App) Tick(ctx context.Context) error {
	start := time.Now()
	if err := ctx.Err(); err != nil {
		return err
	}

	a.ensureStarted()
	a.drain(drainBudget)

	if a.geom {
		a.geom = false
		if err := a.Redraw(ctx); err != nil {
			return err
		}

		a.stats().Redraws++
	} else {
		a.stats().Paints++
	}

	a.paint()
	a.pump(time.Now())
	a.stats().LastTickUS = time.Since(start).Microseconds()

	return nil
}

// ensureStarted asks for the first active set and history page once.
func (a *App) ensureStarted() {
	if a.started {
		return
	}

	a.started = true
	a.needActive = true

	if err := a.askPage(); err != nil {
		a.historyDirty = true
	}
}

// pump refreshes the summary and history at a bounded rate. A request the
// backend cannot accept keeps its flag set, so the next pump retries.
func (a *App) pump(now time.Time) {
	if now.Sub(a.lastSummary) >= summaryEvery {
		a.summaryGen++

		if a.backend.Summary(a.summaryGen) == nil {
			a.lastSummary = now
		}
	}

	if a.needActive && a.backend.Active() == nil {
		a.needActive = false
	}

	if a.historyDirty && now.Sub(a.lastHistory) >= historyEvery {
		if a.askPage() == nil {
			a.historyDirty = false
		}
	}
}

// askPage requests the pager's current page.
func (a *App) askPage() error {
	a.lastHistory = time.Now()

	return a.backend.Page(a.pager.Request())
}

// stats returns the mutable footer counters.
func (a *App) stats() *Stats { return &a.view.Stats }
