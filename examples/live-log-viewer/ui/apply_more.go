package ui

import "time"

// applySources refreshes the sidebar and falls back to all sources when the
// active one was deleted.
func (a *App) applySources(o out) bool {
	if o.gen != a.srcGen.Load() || o.err != nil {
		return false
	}

	a.sources = o.sources

	if a.activeSource != "" && !a.hasSource(a.activeSource) {
		a.activeSource = ""
		a.setNote("the filtered source is gone; showing all sources", time.Now())
		a.loadPage(intentFilter)
	}

	return true
}

// hasSource reports whether the sidebar still lists key.
func (a *App) hasSource(key string) bool {
	for _, s := range a.sources {
		if s.Key == key {
			return true
		}
	}

	return false
}

// applyDetail opens a detail result unless the selection moved on.
func (a *App) applyDetail(o out) bool {
	if o.gen != a.detailGen.Load() || o.id != a.detailID {
		return false
	}

	if o.err != nil {
		a.setNote("detail load failed: "+o.err.Error(), time.Now())

		return true
	}

	a.detail = o.detail
	a.detailOpen = true
	a.scrollTop()

	return true
}

// applySettings loads the persisted view state once.
func (a *App) applySettings(o out) bool {
	if o.err != nil || a.settingsLoaded {
		return false
	}

	a.settingsLoaded = true
	a.dark = o.settings.Dark
	a.activeSevs = o.settings.Severity
	a.activeSource = o.settings.Source
	a.follow.SetFollow(o.settings.Follow)
	a.applyTheme()
	a.loadPage(intentFilter)

	return true
}

// setNote shows a transient status message.
func (a *App) setNote(text string, now time.Time) {
	a.note = text
	a.noteUntil = now.Add(5 * time.Second)
}
