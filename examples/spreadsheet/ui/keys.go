package ui

import "context"

// onKeyDown maps keys to navigation and editing. Typing arrives through
// Type, Backspace, and the Ctrl chords through their own handlers.
func (a *App) onKeyDown(ctx context.Context, key string) error {
	switch key {
	case "shift", "shiftleft", "shiftright":
		a.shift = true

		return nil
	case "control", "controlleft", "controlright", "meta", "metaleft", "metaright":
		a.ctrl = true

		return nil
	case "alt", "altleft", "altright":
		a.alt = true

		return nil
	}

	if a.csv.open {
		return a.csvKey(ctx, key)
	}

	if key == "escape" {
		return a.escapeKey(ctx)
	}

	if a.edit != nil {
		return a.editKey(ctx, key)
	}

	if a.ctrl {
		return a.ctrlKey(ctx, key)
	}

	switch key {
	case "delete":
		return a.clearSelection(ctx)
	case "arrowup":
		return a.step(ctx, -1, 0, a.shift)
	case "arrowdown":
		return a.step(ctx, 1, 0, a.shift)
	case "arrowleft":
		return a.step(ctx, 0, -1, a.shift)
	case "arrowright":
		return a.step(ctx, 0, 1, a.shift)
	case "shift+arrowleft":
		return a.step(ctx, 0, -1, true)
	case "shift+arrowright":
		return a.step(ctx, 0, 1, true)
	case "tab":
		if a.shift {
			return a.step(ctx, 0, -1, false)
		}

		return a.step(ctx, 0, 1, false)
	case "home":
		return a.jump(ctx, 0, a.selection().ActiveR)
	case "end":
		return a.jump(ctx, a.sheet().Cols-1, a.selection().ActiveR)
	case "pageup":
		return a.step(ctx, -a.pageRows(), 0, a.shift)
	case "pagedown":
		return a.step(ctx, a.pageRows(), 0, a.shift)
	case "ctrl+home":
		return a.jump(ctx, 0, 0)
	case "ctrl+end":
		return a.jumpUsed(ctx)
	case "ctrl+arrowleft":
		return a.jump(ctx, 0, a.selection().ActiveR)
	case "ctrl+arrowright":
		return a.jump(ctx, a.sheet().Cols-1, a.selection().ActiveR)
	case "f2":
		s := a.selection()
		a.startEdit(s.ActiveR, s.ActiveC, a.editText())

		return a.Redraw(ctx)
	}

	return nil
}
