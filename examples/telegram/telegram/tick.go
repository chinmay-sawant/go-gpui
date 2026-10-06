package telegram

import "context"

// Tick drains cross-thread requests: a back press, new system bar insets,
// and a captured photo. It redraws when any of them changed the view. The
// window calls it once per frame on the phone.
func (a *App) Tick(ctx context.Context) error {
	a.applyBack()
	a.applyInsets()
	a.applyPhoto()

	if !a.dirty {
		return nil
	}

	a.dirty = false
	a.rebuild()
	a.page.SetData(&a.view)

	return a.page.Redraw(ctx)
}

// mark records that the view needs a rebuild on the next tick.
func (a *App) mark() { a.dirty = true }
