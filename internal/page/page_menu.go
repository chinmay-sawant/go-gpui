package page

import "github.com/chinmay-sawant/go-gpui/internal/host"

// ContextMenu returns the context-menu rows for the current state. Cut and
// copy need a selection, paste needs a focused editable field, select all
// needs either, and undo and redo need their handler.
func (p *Page) ContextMenu() []host.MenuItem {
	sel := p.hasSelection()
	_, editable := p.typingTarget()

	return []host.MenuItem{
		{ID: "cut", Label: "Cut", Enabled: sel},
		{ID: "copy", Label: "Copy", Enabled: sel},
		{ID: "paste", Label: "Paste", Enabled: editable},
		{ID: "select-all", Label: "Select all", Enabled: editable || p.handlers.SelectAll != nil},
		{ID: "undo", Label: "Undo", Enabled: p.handlers.Undo != nil},
		{ID: "redo", Label: "Redo", Enabled: p.handlers.Redo != nil},
	}
}
