package login

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

func (a *App) onClick(_ context.Context, box gpui.Box) error {
	if box.Action == "login" && a.ready() {
		a.signIn()
	}

	a.page.SetData(a.view)

	return nil
}

// onChange refreshes the button state after a field changed.
func (a *App) onChange(context.Context, gpui.Box) error {
	a.refresh()

	return nil
}

// onBeforeEdit snapshots the fields for undo and clears any old message.
func (a *App) onBeforeEdit(context.Context, gpui.Box) error {
	a.push()

	if a.view.Status == "" && a.view.Error == "" {
		return nil
	}

	a.view.Status = ""
	a.view.Error = ""
	a.page.SetData(a.view)

	return nil
}
