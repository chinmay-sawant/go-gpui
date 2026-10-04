package window

// devTab names the panel's top-level views.
type devTab int

const (
	devTabElements devTab = iota
	devTabFrame
	devTabOps
)

// devTabNames are the labels in the tab strip, in order.
var devTabNames = [...]string{"Elements", "Frame", "Ops"}

// devTabAt clamps an index to a known tab.
func devTabAt(i int) devTab {
	if i < 0 {
		return devTabElements
	}

	if i >= len(devTabNames) {
		return devTabOps
	}

	return devTab(i)
}

// devTabHint is the footer hint for one tab.
func devTabHint(tab devTab) string {
	switch tab {
	case devTabFrame:
		return "frame and pipeline counters"
	case devTabOps:
		return "o toggles the outlines"
	default:
		return "click a box in the page to pin it"
	}
}
