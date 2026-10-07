package ui

import (
	"context"
	"unicode/utf8"
)

// ctrlKey handles the modifier chords the page does not own.
func (a *App) ctrlKey(ctx context.Context, key string) error {
	switch key {
	case "arrowleft", "ctrl+arrowleft":
		return a.jump(ctx, 0, a.selection().ActiveR)
	case "arrowright", "ctrl+arrowright":
		return a.jump(ctx, a.sheet().Cols-1, a.selection().ActiveR)
	case "arrowup":
		return a.jump(ctx, a.selection().ActiveC, 0)
	case "arrowdown":
		return a.jump(ctx, a.selection().ActiveC, a.sheet().Rows-1)
	case "home", "ctrl+home":
		return a.jump(ctx, 0, 0)
	case "end", "ctrl+end":
		return a.jump(ctx, a.sheet().Cols-1, a.sheet().Rows-1)
	}

	return nil
}

// editKey handles the keys the editor itself uses.
func (a *App) editKey(ctx context.Context, key string) error {
	e := a.edit

	switch key {
	case "arrowleft", "shift+arrowleft":
		e.moveLeft()
	case "arrowright", "shift+arrowright":
		e.moveRight()
	case "home":
		e.Caret = 0
	case "end":
		e.Caret = utf8.RuneCountInString(e.Text)
	case "tab":
		a.commitEdit(ctx)

		return a.step(ctx, 0, 1, false)
	default:
		return nil
	}

	return a.Redraw(ctx)
}

// escapeKey cancels the editor, then collapses the selection.
func (a *App) escapeKey(ctx context.Context) error {
	if a.edit != nil {
		a.cancelEdit()
	} else {
		s := a.selection()
		a.setSelection(newSelection(s.ActiveR, s.ActiveC))
	}

	return a.Redraw(ctx)
}

// pageRows is the number of visible rows.
func (a *App) pageRows() int {
	return max(1, (a.viewH-ChromeH)/RowH)
}
