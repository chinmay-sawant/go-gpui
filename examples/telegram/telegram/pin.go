package telegram

// composeH is the composer height and sheetH the attachment sheet height,
// both in CSS pixels. The thread reserves them, and Pin keeps the bars at
// the viewport edges while the messages scroll. The top bar height lives in
// components/thread.css.
const (
	composeH = 63
	sheetH   = 150
)

// Pin positions the pinned thread bars for a scroll offset and a viewport
// height. The page's scroll-window callback runs it before each redraw, so
// the window's own translation lands them back at the viewport edges. It
// reports whether the positions changed.
func (a *App) Pin(offsetY, viewH int) bool {
	// The engine places an absolute box against its parent's content box,
	// which starts at InsetTop. The bar must land at offsetY + InsetTop in
	// document space so the window's translation puts it at the top edge.
	top := offsetY
	height := composeH
	if a.view.AttachOpen {
		height += sheetH
	}

	compose := offsetY + viewH - a.view.InsetBottom - height - a.view.InsetTop
	if top == a.view.BarTop && compose == a.view.BottomTop {
		return false
	}

	a.view.BarTop, a.view.BottomTop = top, compose

	return true
}
