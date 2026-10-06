package ui

import "context"

// setPageCancel replaces the in-flight page cancel.
func (a *App) setPageCancel(cancel context.CancelFunc) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.pageCancel != nil {
		a.pageCancel()
	}

	a.pageCancel = cancel
}

// refreshSources asks for the sidebar snapshot.
func (a *App) refreshSources() {
	a.srcGen.Add(1)
	a.send(req{kind: reqSources, gen: a.srcGen.Load(), ctx: a.ctx})
}

// loadSettings asks for the persisted view state once at startup.
func (a *App) loadSettings() {
	a.send(req{kind: reqSettings, ctx: a.ctx})
}

// requestDetail asks for one entry's full message.
func (a *App) requestDetail(id int64) {
	a.detailGen.Add(1)
	a.detailID = id

	ctx, cancel := context.WithCancel(a.ctx)

	a.mu.Lock()
	if a.detailCancel != nil {
		a.detailCancel()
	}

	a.detailCancel = cancel
	a.mu.Unlock()

	a.send(req{kind: reqDetail, gen: a.detailGen.Load(), ctx: ctx, id: id})
}

// saveSettings persists the durable view state.
func (a *App) saveSettings() {
	a.send(req{kind: reqSave, ctx: a.ctx, settings: a.currentSettings()})
}

// currentSettings collects the persistent slice of the view state.
func (a *App) currentSettings() Settings {
	return Settings{
		Dark:     a.dark,
		Follow:   a.follow.Follow,
		Severity: append([]string(nil), a.activeSevs...),
		Source:   a.activeSource,
	}
}

// freeze pins the high-water mark so a history browse ignores new inserts.
func (a *App) freeze() {
	if a.pager.HWM != 0 {
		return
	}

	a.pager.HWM = a.follow.LastSeen
	if n := a.pager.Len(); n > 0 && a.pager.Entries[n-1].ID > a.pager.HWM {
		a.pager.HWM = a.pager.Entries[n-1].ID
	}

	if a.pager.HWM == 0 {
		a.pager.HWM = 1
	}
}
