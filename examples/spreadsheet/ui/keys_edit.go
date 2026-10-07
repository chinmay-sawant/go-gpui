package ui

import "context"

// onType starts an edit with the typed text, or extends the open editor.
// The page draws after this handler runs.
func (a *App) onType(_ context.Context, text string) error {
	if a.csv.open || text == "" {
		return nil
	}

	if a.edit != nil {
		a.edit.insert(text)

		return nil
	}

	s := a.selection()
	a.startEdit(s.ActiveR, s.ActiveC, text)

	return nil
}

// onBackspace edits the open buffer, or opens one with an empty value.
func (a *App) onBackspace(_ context.Context) error {
	if a.csv.open {
		return nil
	}

	if a.edit != nil {
		a.edit.backspace()

		return nil
	}

	s := a.selection()
	a.startEdit(s.ActiveR, s.ActiveC, "")

	return nil
}

// onSubmit commits the editor and moves down, or moves down alone.
func (a *App) onSubmit(ctx context.Context) error {
	if a.csv.open {
		return nil
	}

	a.commitEdit(ctx)

	return a.step(ctx, 1, 0, false)
}

// csvKey handles keys while the CSV dialog is open. The dialog fields use
// the page's own form editing.
func (a *App) csvKey(ctx context.Context, key string) error {
	if key == "escape" {
		return a.closeCSV(ctx)
	}

	return nil
}
