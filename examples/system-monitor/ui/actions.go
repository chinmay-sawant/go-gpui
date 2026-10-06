package ui

import (
	"context"
	"strings"
)

// Action names used in data-action attributes.
const (
	actOverview  = "nav-overview"
	actProcesses = "nav-processes"
	actDetail    = "nav-detail"
	actTheme     = "toggle-theme"
	actMode      = "toggle-mode"
	actRefresh   = "proc-refresh"
	actPrev      = "proc-prev"
	actNext      = "proc-next"
	prefixSort   = "sort-"
	prefixOpen   = "open-"
)

// handleClick routes one clicked action, then rebuilds the template data
// the automatic Redraw paints.
func (a *App) handleClick(ctx context.Context, action string) error {
	switch {
	case action == actOverview:
		a.state.nav = "overview"
	case action == actProcesses:
		a.state.nav = "processes"
	case action == actDetail:
		a.state.nav = "detail"
	case action == actTheme:
		return a.toggleTheme(ctx)
	case action == actMode:
		return a.toggleMode(ctx)
	case action == actRefresh:
		a.refresh()
	case action == actPrev:
		a.state.table.prev()
	case action == actNext:
		a.state.table.next()
	case strings.HasPrefix(action, prefixSort):
		a.state.table.sortBy(sortKey(strings.TrimPrefix(action, prefixSort)))
	case strings.HasPrefix(action, prefixOpen):
		a.pick(strings.TrimPrefix(action, prefixOpen))
	}

	a.sync()

	return nil
}
