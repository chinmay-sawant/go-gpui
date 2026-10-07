package ui

import "time"

// filterState debounces text edits and stamps every filter change with a
// generation. A worker result is accepted only while its generation is
// still current, so a slow old query cannot overwrite a newer filter.
type filterState struct {
	Gen     uint64
	Text    string
	Pending string
	Due     time.Time
	Waiting bool
}

// Edit records a text edit and starts the debounce clock.
func (f *filterState) Edit(text string, now time.Time, delay time.Duration) {
	if f.Waiting && text == f.Pending {
		return
	}

	f.Pending = text
	f.Due = now.Add(delay)
	f.Waiting = true
}

// Ready reports a due edit whose text differs from the applied filter. It
// bumps the generation and returns it with the text to search.
func (f *filterState) Ready(now time.Time) (string, uint64, bool) {
	if !f.Waiting {
		return "", 0, false
	}

	if now.Before(f.Due) {
		return "", 0, false
	}

	f.Waiting = false
	if f.Pending == f.Text {
		return "", 0, false
	}

	f.Text = f.Pending
	f.Gen++

	return f.Text, f.Gen, true
}

// Apply switches the filter immediately, such as a severity or source
// click, and bumps the generation so results in flight are discarded.
func (f *filterState) Apply(text string) uint64 {
	f.Waiting = false
	f.Pending = text
	f.Text = text
	f.Gen++

	return f.Gen
}

// Stale reports that gen belongs to a superseded filter.
func (f *filterState) Stale(gen uint64) bool { return gen != f.Gen }
