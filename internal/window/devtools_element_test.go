package window

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

// TestDevElementRowsJSON checks the pinned box renders as valid JSON with
// the element fields.
func TestDevElementRowsJSON(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.display = devTestDisplay()
	s.dev.havePin, s.dev.pinned = true, d.boxes[0]

	joined := devRowsText(s.devElementRows())
	if !strings.Contains(joined, `"tag": "div"`) || !strings.Contains(joined, `"id": "known"`) {
		t.Fatalf("element rows = %q", joined)
	}

	start := strings.Index(joined, "{")
	end := strings.LastIndex(joined, "}") + 1
	if start < 0 || end <= start || !json.Valid([]byte(joined[start:end])) {
		t.Fatalf("element JSON is not valid: %q", joined)
	}
}

// TestDevElementRowsTarget checks the pinned box wins over the hovered one
// and an empty pick explains how to select.
func TestDevElementRowsTarget(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)

	if text := devRowsText(s.devElementRows()); !strings.Contains(text, "Select an element") {
		t.Fatalf("empty rows = %q", text)
	}

	s.dev.haveHov, s.dev.hovered = true, layout.Box{Tag: "span", ID: "hover"}
	if text := devRowsText(s.devElementRows()); !strings.Contains(text, "hovering  span#hover") {
		t.Fatalf("hover rows = %q", text)
	}

	s.dev.havePin, s.dev.pinned = true, layout.Box{Tag: "div", ID: "pin"}
	if text := devRowsText(s.devElementRows()); !strings.Contains(text, "pinned  div#pin") {
		t.Fatalf("pin rows = %q", text)
	}
}
