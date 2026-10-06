package ui

// visibleRange returns the first and last-plus-one row indexes to render for
// a scroll offset, a viewport height, and an overscan in rows. listTop is
// the y of the first row; rows above it are pinned chrome.
func visibleRange(total, offsetY, viewH, listTop, rowH, over int) (int, int) {
	if total <= 0 || rowH <= 0 {
		return 0, 0
	}

	window := viewH/rowH + 2*over + 1

	start := (offsetY - listTop) / rowH
	if start < 0 {
		start = 0
	}

	start -= over
	if start < 0 {
		start = 0
	}

	end := start + window
	if end > total {
		end = total
		start = end - window
		if start < 0 {
			start = 0
		}
	}

	return start, end
}

// padTop is the laid-out height of the rows above the window.
func padTop(start int, rowH int) int { return start * rowH }

// padBottom is the laid-out height of the rows below the window.
func padBottom(total, end, rowH int) int {
	if end >= total {
		return 0
	}

	return (total - end) * rowH
}
