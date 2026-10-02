package login

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

func (a *App) onClick(_ context.Context, box gpui.Box) error {
	if box.Action == "login" {
		a.signIn()
	}

	a.page.SetData(a.view)

	return nil
}

// beforeType runs before the library appends text to the focused control.
func (a *App) beforeType(_ context.Context, text string) error {
	if text == "" || a.page.FocusedField() == "" {
		return nil
	}

	a.beforeEdit()

	return nil
}

// beforeKey runs before the library backspaces or deletes a word.
func (a *App) beforeKey(context.Context) error {
	if a.page.FocusedField() == "" {
		return nil
	}

	a.beforeEdit()

	return nil
}

// beforeEdit snapshots the fields for undo and clears any old message.
func (a *App) beforeEdit() {
	a.push()

	if a.view.Status == "" && a.view.Error == "" {
		return
	}

	a.view.Status = ""
	a.view.Error = ""
	a.page.SetData(a.view)
}
