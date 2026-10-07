package ui

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

// pageHTML is the one template the window displays.
//
//go:embed ui.html
var pageHTML string

// handlers wires the input callbacks. Handlers that report a command run
// the page's own redraw after them; KeyDown and KeyUp do not, so they call
// Redraw themselves.
func (a *App) handlers() ownframe.Handlers {
	return ownframe.Handlers{
		Click:     a.onClick,
		KeyDown:   a.onKeyDown,
		KeyUp:     a.onKeyUp,
		Type:      a.onType,
		Backspace: a.onBackspace,
		Submit:    a.onSubmit,
		Copy:      a.onCopy,
		Cut:       a.onCut,
		Paste:     a.onPaste,
		SelectAll: a.onSelectAll,
		Undo:      a.onUndo,
		Redo:      a.onRedo,
		Change:    a.onChange,
	}
}
