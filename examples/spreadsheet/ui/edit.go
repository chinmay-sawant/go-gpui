package ui

import "unicode/utf8"

// editing is the text buffer of the active cell editor. The buffer lives in
// the app, not in a form control, so it survives a viewport replacement.
type editing struct {
	Row   int
	Col   int
	Text  string
	Caret int
}

// insert adds text at the caret and moves the caret past it.
func (e *editing) insert(s string) {
	n := utf8.RuneCountInString(e.Text)
	at := clampInt(e.Caret, 0, n)
	runes := []rune(e.Text)
	e.Text = string(runes[:at]) + s + string(runes[at:])
	e.Caret = at + utf8.RuneCountInString(s)
}

// backspace deletes the rune before the caret. It reports whether anything
// changed.
func (e *editing) backspace() bool {
	at := clampInt(e.Caret, 0, utf8.RuneCountInString(e.Text))
	if at == 0 {
		return false
	}

	runes := []rune(e.Text)
	e.Text = string(runes[:at-1]) + string(runes[at:])
	e.Caret = at - 1

	return true
}

// moveLeft and moveRight move the caret one rune without deleting.
func (e *editing) moveLeft() {
	e.Caret = clampInt(e.Caret-1, 0, utf8.RuneCountInString(e.Text))
}

func (e *editing) moveRight() {
	e.Caret = clampInt(e.Caret+1, 0, utf8.RuneCountInString(e.Text))
}

// startEdit opens the editor on a cell with the given text.
func (a *App) startEdit(r, c int, text string) {
	a.setSelection(newSelection(r, c))
	a.edit = &editing{Row: r, Col: c, Text: text, Caret: utf8.RuneCountInString(text)}
	a.status = "edit " + ref(r, c)
}
