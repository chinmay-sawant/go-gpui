package host

// MenuItem is one row of the window's context menu. ID names the action:
// cut, copy, paste, select-all, undo, or redo.
type MenuItem struct {
	ID      string
	Label   string
	Enabled bool
}

// ContextMenu is a screen that reports the context menu rows for its current
// state.
type ContextMenu interface {
	ContextMenu() []MenuItem
}
