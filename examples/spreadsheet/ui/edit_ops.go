package ui

import "context"

// editText returns the raw source of the active cell, or the text already
// in the editor.
func (a *App) editText() string {
	if a.edit != nil {
		return a.edit.Text
	}

	s := a.selection()
	cell, _ := a.tiles.get(a.active, s.ActiveR, s.ActiveC)

	return cell.Raw
}

// cancelEdit drops the editor without touching the workbook.
func (a *App) cancelEdit() {
	if a.edit == nil {
		return
	}

	a.edit = nil
	a.status = "canceled"
}

// commitEdit closes the editor and sends the text to the worker. On a save
// failure the editor is restored so no typed text is lost.
func (a *App) commitEdit(_ context.Context) {
	e := a.edit

	a.edit = nil
	if e == nil {
		return
	}

	s := Selection{AnchorR: e.Row, AnchorC: e.Col, ActiveR: e.Row, ActiveC: e.Col}
	a.setSelection(s)
	cell, _ := a.tiles.get(a.active, e.Row, e.Col)
	if cell.Raw == e.Text {
		return
	}

	a.applyOptimistic(a.active, []Edit{{Row: e.Row, Col: e.Col, Raw: e.Text}})
	a.postEdits(a.active, []Edit{{Row: e.Row, Col: e.Col, Raw: e.Text}})
}
