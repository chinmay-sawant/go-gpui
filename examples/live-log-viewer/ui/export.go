package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// startExport queues a bounded CSV export of the selected range, or of the
// loaded page when nothing is selected.
func (a *App) startExport() {
	if a.feed == nil {
		return
	}

	if a.exporting {
		a.setNote("an export is already running", time.Now())

		return
	}

	q, capped := a.exportQuery()
	if q.ToID == 0 && q.FromID == 0 {
		a.setNote("nothing to export yet", time.Now())

		return
	}

	dir := a.exportDir
	if dir == "" {
		dir = os.TempDir()
	}

	name := "live-log-" + time.Now().Format("20060102-150405") + ".csv"
	path := filepath.Join(dir, name)
	a.exporting = true

	note := fmt.Sprintf("exporting up to %d rows", ExportMax)
	if capped {
		note = fmt.Sprintf("selection capped at %d rows", ExportMax)
	}

	a.setNote(note, time.Now())
	a.send(req{kind: reqExport, ctx: a.ctx, q: q, path: path})
}

// exportQuery builds the inclusive ID range to export.
func (a *App) exportQuery() (Query, bool) {
	q := a.baseQuery()
	q.Limit = ExportMax
	q.MaxID = a.pager.HWM

	lo, hi := a.selectionRange()
	if lo == 0 {
		if n := a.pager.Len(); n > 0 {
			lo, hi = a.pager.Entries[0].ID, a.pager.Entries[n-1].ID
		}
	}

	q.FromID, q.ToID = lo, hi

	return q, hi-lo+1 > ExportMax
}

// selectionRange returns the inclusive range between the selection anchor
// and the selected row, or zero when they are equal.
func (a *App) selectionRange() (int64, int64) {
	if a.selAnchor == 0 || a.selected == 0 || a.selAnchor == a.selected {
		return 0, 0
	}

	if a.selAnchor > a.selected {
		return a.selected, a.selAnchor
	}

	return a.selAnchor, a.selected
}

// applyExport reports the export result.
func (a *App) applyExport(o out, now time.Time) bool {
	a.exporting = false

	if o.err != nil {
		a.setNote("export failed: "+o.err.Error(), now)

		return true
	}

	a.setNote(fmt.Sprintf("exported %d rows to %s", o.count, o.path), now)

	return true
}
