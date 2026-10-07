package ui

import "strings"

// applyOptimistic paints an edit before the backend confirms it.
func (a *App) applyOptimistic(sheet string, edits []Edit) {
	for _, ed := range edits {
		cell, _ := a.tiles.get(sheet, ed.Row, ed.Col)
		next := Cell{Raw: ed.Raw, Display: ed.Raw}
		if strings.HasPrefix(ed.Raw, "=") {
			next.Display = cell.Display
		}

		a.tiles.setCell(sheet, ed.Row, ed.Col, next)
	}
}

// postEdits queues one atomic edit batch.
func (a *App) postEdits(sheet string, edits []Edit) {
	a.work.post(job{kind: jobApply, sheet: sheet, edits: edits})
}
