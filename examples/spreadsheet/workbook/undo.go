package workbook

// Undo reverts the last command. The bool is false when nothing is left.
func (w *Workbook) Undo() (bool, RecalcResult) {
	if len(w.history) == 0 {
		return false, RecalcResult{}
	}

	h := w.history[len(w.history)-1]
	w.history = w.history[:len(w.history)-1]

	res := w.replay(h.inverse)
	w.future = append(w.future, h)

	return true, res
}

// Redo reapplies the last undone command.
func (w *Workbook) Redo() (bool, RecalcResult) {
	if len(w.future) == 0 {
		return false, RecalcResult{}
	}

	h := w.future[len(w.future)-1]
	w.future = w.future[:len(w.future)-1]

	res := w.replay(h.forward)
	w.history = append(w.history, h)

	return true, res
}

// CanUndo reports whether an undo step is available.
func (w *Workbook) CanUndo() bool { return len(w.history) > 0 }

// CanRedo reports whether a redo step is available.
func (w *Workbook) CanRedo() bool { return len(w.future) > 0 }

// UndoDepth is the number of undo steps held.
func (w *Workbook) UndoDepth() int { return len(w.history) }

// RedoDepth is the number of redo steps held.
func (w *Workbook) RedoDepth() int { return len(w.future) }

// SetHistoryLimit bounds the undo stack. A value below one restores
// DefaultHistory, and neither stack keeps more than the limit.
func (w *Workbook) SetHistoryLimit(n int) {
	if n <= 0 {
		n = DefaultHistory
	}

	w.limit = n

	if len(w.history) > n {
		w.history = w.history[len(w.history)-n:]
	}

	if len(w.future) > n {
		w.future = w.future[len(w.future)-n:]
	}
}
