package ui

import (
	_ "embed"
	"time"
)

// darkCSS is the extra stylesheet for dark mode; the base sheet in ui.html
// is the light theme.
//
//go:embed theme_dark.css
var darkCSS string

// applyTheme installs the stylesheet for the current mode.
func (a *App) applyTheme() error {
	if !a.dark {
		return a.page.SetTheme("")
	}

	return a.page.SetTheme(darkCSS)
}

// toggleTheme flips light and dark and persists the choice.
func (a *App) toggleTheme() {
	a.dark = !a.dark

	if err := a.applyTheme(); err != nil {
		a.setNote("theme sheet rejected: "+err.Error(), time.Now())
	}

	a.saveSettings()
}
