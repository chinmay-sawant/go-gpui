package workbook

// SavedRev is the storage revision the workbook was last acknowledged at:
// the value read on load and the value SaveCells returned after each
// commit. Callers pass it as the base revision for the next save so a
// second window's write is detected as a conflict.
func (w *Workbook) SavedRev() int64 { return w.saved }

// SetSavedRev records an acknowledged save.
func (w *Workbook) SetSavedRev(rev int64) { w.saved = rev }

// Dirty reports whether edits happened since the last acknowledged save.
func (w *Workbook) Dirty() bool { return w.rev != w.saved }
