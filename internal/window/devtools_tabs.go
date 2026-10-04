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
