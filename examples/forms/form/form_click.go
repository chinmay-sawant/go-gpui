package form

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
)

func (a *App) onClick(_ context.Context, box gpui.Box) error {
	if box.Action != "send" {
		return nil
	}

	a.view.Status = fmt.Sprintf(
		"email=%s remember=%t plan=%s color=%s file=%s",
		a.page.FormValue("email"),
		a.page.FormChecked("remember"),
		a.plan(),
		a.page.FormValue("color"),
		a.page.FormValue("file"),
	)
	a.page.SetData(a.view)

	return nil
}

func (a *App) plan() string {
	if a.page.FormChecked("pro") {
		return a.page.FormValue("pro")
	}

	if a.page.FormChecked("free") {
		return a.page.FormValue("free")
	}

	return ""
}
