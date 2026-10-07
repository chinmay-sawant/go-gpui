package ui

import "context"

// commitCSV applies the previewed bytes atomically.
func (a *App) commitCSV(ctx context.Context) error {
	if len(a.csv.data) == 0 {
		a.status = "import: choose a file first"

		return a.Redraw(ctx)
	}

	a.status = "importing"
	a.work.post(job{kind: jobCommitCSV, sheet: a.sheet().ID, data: a.csv.data, flag: a.csv.replace})

	return a.Redraw(ctx)
}

// writeCSV exports the active sheet to the dialog path.
func (a *App) writeCSV(ctx context.Context) error {
	if a.csv.path == "" {
		a.status = "export: empty path"

		return a.Redraw(ctx)
	}

	a.status = "exporting to " + a.csv.path
	a.work.post(job{kind: jobExport, sheet: a.sheet().ID, path: a.csv.path})

	return a.Redraw(ctx)
}

// previewData requests a preview of already-read bytes.
func (a *App) previewData() {
	a.work.post(job{kind: jobPreview, sheet: a.sheet().ID, data: a.csv.data})
}
