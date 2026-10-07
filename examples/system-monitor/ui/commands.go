package ui

import "context"

// toggleTheme switches the extra sheet and persists the choice. The click's
// automatic Redraw applies the sheet.
func (a *App) toggleTheme(ctx context.Context) error {
	dark := !a.state.dark

	if err := a.page.SetTheme(themeSource(dark)); err != nil {
		a.state.notice = "theme failed: " + err.Error()

		return nil
	}

	a.state.dark = dark
	a.state.notice = ""
	a.saveSettings(ctx)

	return nil
}

// toggleMode switches the source and resets every baseline the UI owns, so
// live data never mixes with dummy history.
func (a *App) toggleMode(ctx context.Context) error {
	live := !a.state.live

	if err := a.src.SetLive(live); err != nil {
		a.state.notice = "cannot switch to " + modeName(live) + ": " + err.Error()

		return nil
	}

	a.state.reset()
	a.state.live = live
	a.state.mode = modeName(live)
	a.state.notice = "switched to " + a.state.mode + " mode"
	a.src.Track("")

	return nil
}

// refresh adopts the pending snapshot and keeps the selection by identity,
// even when the process arrived or exited between snapshots.
func (a *App) refresh() {
	if !a.state.table.refresh() {
		a.state.notice = "no new sample yet"

		return
	}

	a.state.notice = ""
	sel := &a.state.sel

	if sel.id != "" {
		if p, ok := a.state.table.find(sel.id); ok {
			sel.proc = p
			sel.found = true
		} else {
			sel.found = false
		}
	}
}

// pick selects a process identity for the detail view. An empty id clears
// the selection; a stale result never attaches because drain compares ids.
func (a *App) pick(id string) {
	sel := &a.state.sel
	*sel = selection{id: id}

	if id == "" {
		a.src.Track("")

		return
	}

	sel.trk = Tracked{ID: id, Loading: true}

	if p, ok := a.state.table.find(id); ok {
		sel.proc = p
		sel.found = true
	}

	a.src.Track(id)
}
