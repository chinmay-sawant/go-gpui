package telegram

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick runs the control under the tap and rebuilds the view.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	switch {
	case box.Action == "tab-chats":
		a.view.Tab = "chats"
	case box.Action == "tab-contacts":
		a.view.Tab = "contacts"
	case box.Action == "tab-settings":
		a.view.Tab = "settings"
	case box.Action == "chat-back":
		a.view.Active = ""
		a.view.Status = ""
	case strings.HasPrefix(box.Action, "open-"):
		a.open(strings.TrimPrefix(box.Action, "open-"))
	case strings.HasPrefix(box.Action, "contact-"):
		a.openContact(strings.TrimPrefix(box.Action, "contact-"))
	case box.Action == "send":
		a.send()
	default:
		return nil
	}

	a.rebuild()
	a.page.SetData(&a.view)

	return nil
}

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
