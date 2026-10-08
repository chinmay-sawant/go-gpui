package remote

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/ownframe"
)

func (a *App) onClick(_ context.Context, box ownframe.Box) error {
	a.host = strings.TrimSpace(a.page.FormValue("host"))

	switch box.ID {
	case "theme":
		a.toggleTheme()
	case "mode":
		a.toggleMode()
	case "connect":
		a.connect()
	case "scan":
		a.view.Status = "Searching this Wi-Fi"
		a.later(a.doFind)
	case "":
	default:
		a.press(box.ID)
	}

	a.page.SetData(&a.view)

	return nil
}

func (a *App) onKey(ctx context.Context, key string) error {
	if a.page.FocusedField() != "" {
		return nil
	}

	id := map[string]string{
		"arrowup": "up", "arrowdown": "down",
		"arrowleft": "left", "arrowright": "right",
		"enter": "ok", "escape": "back",
	}[key]
	if id == "" {
		return nil
	}

	a.host = strings.TrimSpace(a.page.FormValue("host"))
	a.press(id)
	a.page.SetData(&a.view)

	return a.page.Redraw(ctx)
}

func (a *App) toggleTheme() {
	a.dark = !a.dark
	css := lightTheme
	a.view.ThemeLabel = "Dark"
	a.view.Status = "Light theme"
	if a.dark {
		css = darkTheme
		a.view.ThemeLabel = "Light"
		a.view.Status = "Dark theme"
	}

	if err := a.page.SetTheme(css); err != nil {
		a.view.Status = err.Error()
	}
}
