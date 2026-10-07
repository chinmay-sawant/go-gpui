package ui

import "fmt"

// applyCSV folds a file read, preview, commit, or export answer into the
// dialog state.
func (a *App) applyCSV(r result) bool {
	switch r.kind {
	case jobReadFile:
		if r.err != nil {
			a.status = "import: " + r.err.Error()
			a.csv.note = a.status

			return true
		}

		a.csv.data = r.data
		a.csv.note = "parsing"
		a.previewData()

		return true
	case jobPreview:
		if r.err != nil {
			a.status = "import: " + r.err.Error()
			a.csv.note = a.status

			return true
		}

		a.csv.prev = r.prev
		a.csv.ready = true
		a.csv.note = fmt.Sprintf("%d rows x %d cols", r.prev.Total, r.prev.Cols)

		return true
	case jobCommitCSV:
		if r.err != nil {
			a.status = "import failed: " + r.err.Error()
			a.csv.note = "workbook unchanged"

			return true
		}

		a.onRev(r.imp.Rev)
		a.csv = csvState{}
		a.status = fmt.Sprintf("imported %d rows", r.imp.Rows)

		return true
	case jobExport:
		if r.err != nil {
			a.status = "export failed: " + r.err.Error()

			return true
		}

		a.csv = csvState{}
		a.status = "wrote " + r.path

		return true
	}

	return false
}
