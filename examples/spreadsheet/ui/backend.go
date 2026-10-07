package ui

// Backend is the workbook engine the screen reads and edits. The core
// workbook, formula, and storage packages implement it; tests use a fake.
// Every method may run on the worker goroutine, so an implementation must
// be safe for one caller at a time.
type Backend interface {
	Sheets() []Sheet
	Info(id string) (Sheet, bool)
	Revision() uint64
	// Range returns the cells of area in row-major order, exactly
	// area.W()*area.H() entries.
	Range(id string, area Area) ([]Cell, error)
	// Used returns the used rectangle of a sheet, or false when it is empty.
	Used(id string) (Area, bool)
	// Apply stores the edits and returns the new revision.
	Apply(id string, edits []Edit) (uint64, error)
	Undo() (UndoResult, error)
	Redo() (UndoResult, error)
	Pref(key string) (string, bool)
	SetPref(key, value string) error
	PreviewCSV(id string, data []byte, replace bool) (Preview, error)
	CommitCSV(id string, data []byte, replace bool) (ImportResult, error)
	ExportCSV(id string) (string, error)
	Close() error
}

// Sheet is one sheet's identity and size.
type Sheet struct {
	ID   string
	Name string
	Rows int
	Cols int
}

// Cell is one sheet cell as the backend reports it. Raw is the stored
// source, Display the painted value, Num marks a number, and Err carries a
// formula error to paint.
type Cell struct {
	Raw     string
	Display string
	Num     bool
	Err     string
}

// Edit replaces the raw source of one cell.
type Edit struct {
	Row int
	Col int
	Raw string
}

// UndoResult reports one applied history step.
type UndoResult struct {
	Rev   uint64
	OK    bool
	Label string
}

// Preview is a parsed CSV import before it is committed.
type Preview struct {
	Rows     [][]string
	Total    int
	Cols     int
	Warnings []string
}

// ImportResult reports a committed CSV import.
type ImportResult struct {
	Rev  uint64
	Rows int
	Cols int
}
