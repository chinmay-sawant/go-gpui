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
	a.historyDirty = false
	a.needActive = true
	a.askPage()
}

// pump refreshes the summary and history at a bounded rate.
func (a *App) pump(now time.Time) {
	if now.Sub(a.lastSummary) >= summaryEvery {
		a.lastSummary = now
		a.summaryGen++
		a.backend.Summary(a.summaryGen)
	}

	if a.needActive {
		a.needActive = false
		a.backend.Active()
	}

	if a.historyDirty && now.Sub(a.lastHistory) >= historyEvery {
		a.historyDirty = false
		a.askPage()
	}
}

// askPage requests the pager's current page.
func (a *App) askPage() {
	a.lastHistory = time.Now()
	a.backend.Page(a.pager.Request())
}

// stats returns the mutable footer counters.
func (a *App) stats() *Stats { return &a.view.Stats }
