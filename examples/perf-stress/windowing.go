package main

// rowH is the fixed grid row height in CSS pixels: 26px content with 5px
// padding top and bottom. The windowing test pins it to a laid-out box,
// so a restyle that changes row geometry fails loudly instead of
// silently misaligning the row window.
const rowH = 36

// winOver is the overscan in rows above and below the viewport. Scrolling
// inside it reuses the laid-out window; crossing it slices and redraws.
const winOver = 10

// gridTop finds the grid origin from the last layout, or zero before the
// first draw. The header above it has fixed CSS, so it stays put.
func (a *App) gridTop() float64 {
	for _, b := range a.page.Boxes() {
		if b.ID == "grid" {
			return b.Y
		}
	}

	return 0
}

// applyWindow slices the visible rows plus overscan for a scroll offset.
// It calls SetData only when the window moved, so scrolling inside one
// window redraws nothing and crossing it pays one small redraw.
func (a *App) applyWindow(offsetY, viewH int) {
	total := len(a.all)
	if total == 0 {
		return
	}

	start := (offsetY-int(a.gridTop()))/rowH - winOver
	if start < 0 {
		start = 0
	}

	end := start + viewH/rowH + 2*winOver + 1
	if end > total {
		end = total
		start = end - (viewH/rowH + 2*winOver + 1)
		if start < 0 {
			start = 0
		}
	}

	if start == a.winStart && end == a.winEnd {
		return
	}

	a.winStart, a.winEnd = start, end
	a.view.Rows = a.all[start:end]
	a.view.TopPad = start * rowH
	a.view.BotPad = (total - end) * rowH
	a.page.SetData(a.view)
}

// ScrollToRow queues a scroll that puts row n at the top of the grid.
func (a *App) ScrollToRow(n int) {
	if n < 0 {
		n = 0
	}

	if n >= len(a.all) {
		n = len(a.all) - 1
	}

	a.page.ScrollTo(0, int(a.gridTop())+n*rowH)
}
