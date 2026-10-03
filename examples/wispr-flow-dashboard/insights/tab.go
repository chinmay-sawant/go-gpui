package insights

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// tabNames maps a tab action to the active-tab name.
var tabNames = map[string]string{
	"tab-usage":       "usage",
	"tab-voice":       "voice",
	"tab-leaderboard": "leaderboard",
}

// tabTitles maps a placeholder tab to its heading.
var tabTitles = map[string]string{
	"voice":       "Your voice",
	"leaderboard": "Leaderboard",
}

// onClick switches the active tab.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	name, ok := tabNames[box.Action]
	if !ok {
		return nil
	}

	a.view.ActiveTab = name

	if title, ok := tabTitles[name]; ok {
		a.view.EmptyTitle = title
	}

	a.page.SetData(a.view)

	return nil
}
