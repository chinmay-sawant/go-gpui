package window

import (
	"strings"
	"testing"
)

// devRowsText joins the plain text of rows for assertions.
func devRowsText(rows []devRow) string {
	lines := make([]string, len(rows))

	for i, row := range rows {
		lines[i] = row.line.plain()
	}

	return strings.Join(lines, "\n")
}

// devRowContaining finds the first row whose plain text holds a label.
func devRowContaining(rows []devRow, label string) (devRow, bool) {
	for _, row := range rows {
		if strings.Contains(row.line.plain(), label) {
			return row, true
		}
	}

	return devRow{}, false
}

// devFrameValue returns the trimmed value of the first row whose plain text
// starts with a label, for Frame tab assertions.
func devFrameValue(rows []devRow, label string) string {
	for _, row := range rows {
		plain := row.line.plain()
		if strings.HasPrefix(plain, label+" ") {
			return strings.TrimSpace(strings.TrimPrefix(plain, label))
		}
	}

	return ""
}

// TestDevPanelShowsRelayoutCounts checks the Frame tab carries the relayout
// and skip counters.
func TestDevPanelShowsRelayoutCounts(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.commits, s.skipped = 3, 2

	rows := s.devFrameRows(320)
	row, ok := devRowContaining(rows, "Relayouts")
	if !ok || !strings.HasSuffix(strings.TrimRight(row.line.plain(), " "), "3") {
		t.Fatalf("relayout row = %q, want the value 3", row.line.plain())
	}

	row, ok = devRowContaining(rows, "Skipped")
	if !ok || !strings.HasSuffix(strings.TrimRight(row.line.plain(), " "), "2") {
		t.Fatalf("skipped row = %q, want the value 2", row.line.plain())
	}
}
