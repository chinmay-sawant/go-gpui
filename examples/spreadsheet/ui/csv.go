package ui

import (
	"context"
	"os"
	"path/filepath"
)

// csvState is the import or export dialog.
type csvState struct {
	open    bool
	export  bool
	replace bool
	path    string
	data    []byte
	prev    Preview
	ready   bool
	note    string
}

// closeCSV hides the dialog.
func (a *App) closeCSV(ctx context.Context) error {
	a.csv = csvState{}

	return a.Redraw(ctx)
}

// dataDir is the directory exports are written to by default.
func (a *App) dataDir() string {
	if a.opts.DataDir != "" {
		return a.opts.DataDir
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "."
	}

	return filepath.Join(dir, "ownframe", "spreadsheet")
}

// exportPath is the default export path for the active sheet.
func (a *App) exportPath() string {
	return filepath.Join(a.dataDir(), a.sheet().Name+".csv")
}

// loadCSV reads a chosen file and asks the backend for a preview. The read
// and the parse run on the worker.
func (a *App) loadCSV(path string) {
	a.csv.path = path
	a.csv.ready = false
	a.csv.note = "reading " + path
	a.status = a.csv.note

	a.work.post(job{kind: jobReadFile, path: path, sheet: a.sheet().ID})
}
