package workbook

// histEntry holds both directions of one applied command.
type histEntry struct {
	inverse Command
	forward Command
}

// push records a change and drops the redo stack.
func (w *Workbook) push(h histEntry) {
	w.history = append(w.history, h)

	if n := len(w.history); n > w.limit {
		w.history = w.history[n-w.limit:]
	}

	w.future = nil
}

// replay applies a stored command without touching the history.
func (w *Workbook) replay(cmd Command) RecalcResult {
	s := w.Sheet(cmd.Sheet)
	if s == nil {
		return RecalcResult{}
	}

	for _, e := range cmd.Edits {
		w.putCell(s, e.Pos, e.Cell)
	}

	return w.bump(s, cmd.Edits)
}
