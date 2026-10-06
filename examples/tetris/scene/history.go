package scene

// PageSize is how many score rows one history page carries.
const PageSize = 20

// history is the score-history screen: its open state, the loaded page,
// and whether it shows the seeded demo entries instead of live rankings.
type history struct {
	open    bool
	page    int
	more    bool
	demo    bool
	loading bool
	msg     string
	entries []ScoreEntry
	wasRun  bool
}

// want clamps a page request to the pages known so far.
func (h *history) want(page int) int {
	if page < 0 {
		return 0
	}

	return page
}

// canPrev reports whether a previous page exists.
func (h *history) canPrev() bool { return h.page > 0 }

// canNext reports whether a next page may exist.
func (h *history) canNext() bool { return h.more }

// apply stores one score page.
func (h *history) apply(p ScorePage) {
	h.page = p.Index
	h.more = p.More
	h.entries = p.Entries
	h.loading = false
	h.msg = ""

	if len(p.Entries) == 0 && !p.More {
		if h.demo {
			h.msg = "NO DEMO SCORES"
		} else {
			h.msg = "NO SCORES YET"
		}
	}
}

// rows returns the printable rows, padded to one page with empty runs.
func (h *history) rows() []rowV {
	out := make([]rowV, PageSize)
	for i := range out {
		out[i].Top = histRowY + i*histRowStep

		if i < len(h.entries) {
			out[i].Text = entryLine(h.entries[i])
		}
	}

	return out
}

// title is the history header for the current mode.
func (h *history) title() string {
	if h.demo {
		return "DEMO SCORES"
	}

	return "HIGH SCORES"
}

// demoToggle is the caption of the mode button.
func (h *history) demoToggle() string {
	if h.demo {
		return "LIVE"
	}

	return "DEMO"
}
