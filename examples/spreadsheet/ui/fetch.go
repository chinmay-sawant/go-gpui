package ui

// requestArea posts a bounded fetch. A newer fetch replaces a queued one,
// so stale range loads never run.
func (a *App) requestArea(area Area) {
	if area.Empty() {
		return
	}

	a.fetchGen++
	a.work.post(job{
		kind:  jobFetch,
		gen:   a.fetchGen,
		sheet: a.active,
		area:  area,
	})
}

// requestWindow refetches the current fetch window, after a revision
// change or a sheet switch.
func (a *App) requestWindow() {
	a.requestArea(a.fetchWin.Area())
}
