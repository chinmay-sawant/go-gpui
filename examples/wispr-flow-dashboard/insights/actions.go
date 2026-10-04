package insights

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
	osclip "github.com/chinmay-sawant/go-gpui/internal/clipboard"
)

// shareText is the link the share badge copies.
const shareText = "https://wisprflow.ai/insights"

// onClick runs the control under the click: the streak chevrons, the share
// badge, the mobile download, and the tabs.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	switch box.Action {
	case "streak-prev":
		a.shiftStreak(1)
	case "streak-next":
		a.shiftStreak(-1)
	case "share":
		osclip.Write(shareText)
		a.view.Note = "Share link copied"
	case "download":
		a.view.Note = "Download link sent to your phone"
	}

	if name, ok := tabNames[box.Action]; ok {
		a.view.ActiveTab = name
		a.view.Note = ""

		if title, ok := tabTitles[name]; ok {
			a.view.EmptyTitle = title
		}
	}

	a.page.SetData(a.view)

	return nil
}
