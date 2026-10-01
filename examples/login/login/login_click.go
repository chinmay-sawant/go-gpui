package login

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

func (a *App) onClick(_ context.Context, box gpui.Box) error {
	switch box.Action {
	case "focus":
		if a.view.Focus != box.ID {
			a.view.Selected = false
		}

		a.view.Focus = box.ID
	case "login":
		a.signIn()
	}

	a.page.SetData(a.view)

	return nil
}

func (a *App) onType(_ context.Context, text string) error {
	field := a.focused()
	if field == nil || text == "" {
		return nil
	}

	next := *field + text
	if a.view.Selected {
		next = text
	}

	a.setField(field, next)

	return nil
}

func (a *App) onBackspace(_ context.Context) error {
	field := a.focused()
	if field == nil {
		return nil
	}

	next := ""
	if !a.view.Selected {
		next = dropLastRune(*field)
	}

	a.setField(field, next)

	return nil
}

func (a *App) onDeleteWord(_ context.Context) error {
	field := a.focused()
	if field == nil {
		return nil
	}

	next := ""
	if !a.view.Selected {
		next = dropLastWord(*field)
	}

	a.setField(field, next)

	return nil
}

func (a *App) onPaste(ctx context.Context, text string) error {
	return a.onType(ctx, text)
}
