package ui

import "context"

// loadSettings reads the stored settings, falling back to defaults when the
// store is missing or fails, so a broken store never blocks startup.
func loadSettings(ctx context.Context, st Store) Settings {
	if st == nil {
		return Settings{}
	}

	s, err := st.LoadSettings(ctx)
	if err != nil {
		return Settings{}
	}

	return s
}

// saveSettings persists the theme toggle. A failure becomes a notice; the
// toggle itself still applies to the running window.
func (a *App) saveSettings(ctx context.Context) {
	if a.cfg.Store == nil {
		return
	}

	if err := a.cfg.Store.SaveSettings(ctx, Settings{Dark: a.state.dark}); err != nil {
		a.state.notice = "settings not saved: " + err.Error()
	}
}
