package ui

// tablePage is one bounded window over the filtered snapshot.
type tablePage struct {
	Rows  []Process
	No    int
	Pages int
	Total int
	Shown int
}

// pageView slices the frozen snapshot to the current page. Rows holds at
// most rowsPerPage entries, whatever the total process count.
func (t *table) pageView() tablePage {
	all := t.filtered()
	pages := pageCount(len(all), rowsPerPage)

	no := t.page
	if no > pages {
		no = pages
	}

	if no < 1 {
		no = 1
	}

	start := (no - 1) * rowsPerPage
	end := min(start+rowsPerPage, len(all))

	return tablePage{
		Rows:  all[start:end],
		No:    no,
		Pages: pages,
		Total: len(t.shown.Procs),
		Shown: len(all),
	}
}

// pageCount returns at least one page so an empty list still reads "1 / 1".
func pageCount(n, per int) int {
	if n <= 0 {
		return 1
	}

	return (n + per - 1) / per
}
