package ui

import (
	"context"
	"time"
)

// buildView writes the non-window parts of the template data.
func (a *App) buildView(now time.Time) {
	a.width, _ = a.page.Size()
	a.view.Mode = "list"

	if a.detailOpen {
		a.view.Mode = "detail"
	}

	a.view.Dark = a.dark
	a.view.Sources = sourceRows(a.sources, a.activeSource)
	a.view.Severities = chips(a.minSev)
	a.view.QueryText = a.filters.Text
	a.view.Follow = a.follow.Follow
	a.view.Paused = a.follow.Paused
	a.view.Unread = a.follow.Unread
	a.view.HasOlder = a.pager.HasOlder
	a.view.HasNewer = a.pager.HasNewer
	a.view.Detail = a.detailView()
	a.view.Status = a.statusText(now)
	a.view.Right = a.rightText(now)
	a.view.PageLabel = a.pageLabel()
}

// draw refreshes the template data and renders the page.
func (a *App) draw(ctx context.Context) error {
	a.buildView(time.Now())
	a.page.SetData(&a.view)

	if !a.perf {
		return a.page.Redraw(ctx)
	}

	start := time.Now()
	err := a.page.Redraw(ctx)
	took := time.Since(start)
	a.lastDraw = took

	if took > a.maxDraw {
		a.maxDraw = took
	}

	return err
}

// emptyText explains an empty page.
func (a *App) emptyText() string {
	if a.filters.Text != "" || a.minSev != "" {
		return "No entries match the current filter."
	}

	if a.activeSource != "" {
		return "No entries from this source yet."
	}

	return "No entries yet."
}
