package ui

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/ownframe"
)

// onClick applies the control under the pointer. A non-fatal failure is a
// notice, because a handler error would close the window.
func (a *App) onClick(ctx context.Context, box ownframe.Box) error {
	action := box.Action

	switch {
	case action == "theme":
		a.toggleTheme()
	case action == "add":
		a.addJob()
	case action == "refresh":
		a.refresh()
	case action == "detail-close":
		a.clearDetail()
	case strings.HasPrefix(action, "select-"):
		a.selectRow(strings.TrimPrefix(action, "select-"))
	case strings.HasPrefix(action, "pause-"):
		a.control(ControlPause, strings.TrimPrefix(action, "pause-"))
	case strings.HasPrefix(action, "resume-"):
		a.control(ControlResume, strings.TrimPrefix(action, "resume-"))
	case strings.HasPrefix(action, "cancel-"):
		a.control(ControlCancel, strings.TrimPrefix(action, "cancel-"))
	case strings.HasPrefix(action, "retry-"):
		a.control(ControlRetry, strings.TrimPrefix(action, "retry-"))
	case strings.HasPrefix(action, "remove-"):
		a.control(ControlRemove, strings.TrimPrefix(action, "remove-"))
	case action == "page-next":
		if a.pager.Next() {
			a.askPage()
		}
	case action == "page-prev":
		if a.pager.Prev() {
			a.askPage()
		}
	case action == "page-refresh":
		a.askPage()
	case strings.HasPrefix(action, "filter-"):
		a.setFilter(ParseFilter(strings.TrimPrefix(action, "filter-")))
	}

	a.syncPage()
	a.page.SetData(a.view)

	return nil
}
