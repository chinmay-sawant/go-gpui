package window

import (
	"strings"
	"testing"
)

// TestDevLayoutPlacesTheDock checks the dock hugs the right edge, the
// content rect has room, and every tab gets a hit.
func TestDevLayoutPlacesTheDock(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)

	s.devLayout()

	if s.dev.panel.X != float64(s.screenW)-devDockWidth || s.dev.panel.H != float64(s.screenH) {
		t.Fatalf("panel = %+v", s.dev.panel)
	}

	if s.dev.content.W <= 0 || s.dev.content.H <= 0 {
		t.Fatalf("content = %+v", s.dev.content)
	}

	tabs := 0

	for _, hit := range s.dev.hits {
		if hit.act == devActTab {
			tabs++
		}
	}

	if tabs != len(devTabNames) {
		t.Fatalf("tab hits = %d, want %d", tabs, len(devTabNames))
	}
}

// TestDevLayoutTabContent checks each tab builds its own rows.
func TestDevLayoutTabContent(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)

	s.dev.tab = devTabElements
	if text := devRowsText(s.devTabRows()); !strings.Contains(text, "Select an element") {
		t.Fatalf("elements tab = %q", text)
	}

	s.dev.tab = devTabFrame
	if text := devRowsText(s.devTabRows()); !strings.Contains(text, "PIPELINE") {
		t.Fatalf("frame tab = %q", text)
	}

	s.display = devTestDisplay()
	s.dev.tab = devTabOps
	if text := devRowsText(s.devTabRows()); !strings.Contains(text, "PAINT ORDER") {
		t.Fatalf("ops tab = %q", text)
	}
}
