package telegram

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// onChange rebuilds the list after a search keystroke and swaps the theme
// after the dark-mode toggle.
func (a *App) onChange(_ context.Context, box gpui.Box) error {
	switch box.ID {
	case "search":
		a.rebuild()
	case "dark":
		if err := a.page.SetTheme(themeSource(a.view.Dark)); err != nil {
			return err
		}
	}

	a.page.SetData(&a.view)

	return nil
}

// onSubmit sends the composer on Enter or NumpadEnter.
func (a *App) onSubmit(_ context.Context) error {
	if a.page.FocusID() != "compose" {
		return nil
	}

	a.send()
	a.rebuild()
	a.page.SetData(&a.view)

	return nil
}
