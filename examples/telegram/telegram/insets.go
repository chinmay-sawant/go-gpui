package telegram

// SetInsets records the system bar insets in CSS pixels. The activity calls
// it from the UI thread; the next tick pads the page.
func (a *App) SetInsets(top, bottom int) {
	a.insetTop.Store(int64(top))
	a.insetBot.Store(int64(bottom))
}

// applyInsets pads the page for the system bars and keeps the newest
// message above the keyboard when the bottom inset grows.
func (a *App) applyInsets() {
	top, bottom := int(a.insetTop.Load()), int(a.insetBot.Load())
	if top == a.view.InsetTop && bottom == a.view.InsetBottom {
		return
	}

	if bottom > a.view.InsetBottom {
		a.page.ScrollTo(0, 1<<20)
	}

	a.view.InsetTop, a.view.InsetBottom = top, bottom
	a.mark()
}
