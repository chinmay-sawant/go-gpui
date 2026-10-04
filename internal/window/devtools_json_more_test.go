package window

import (
	"strings"
	"testing"
)

// TestDevJSONCollapse checks a folded node and its click key.
func TestDevJSONCollapse(t *testing.T) {
	t.Parallel()

	rows := devJSONRows(devJSONFixture(), map[string]bool{"rect": true})
	if text := devRowsText(rows); !strings.Contains(text, `"rect": {...}`) {
		t.Fatalf("collapsed text = %q", text)
	}

	if !devJSONHasAct(rows, "rect") {
		t.Fatal("the folded row carries no devActJSON hit")
	}
}

// TestDevJSONArrayPath checks an element path uses its index.
func TestDevJSONArrayPath(t *testing.T) {
	t.Parallel()

	rows := devJSONRows(devJSONFixture(), map[string]bool{"items.0": true})
	if !devJSONHasAct(rows, "items.0") {
		t.Fatalf("no fold row for items.0 in %q", devRowsText(rows))
	}
}

// TestDevJSONRootUnfoldable checks the root opener does not fold.
func TestDevJSONRootUnfoldable(t *testing.T) {
	t.Parallel()

	for _, row := range devJSONRows(devJSONFixture(), nil) {
		if row.act == devActJSON && row.key == "" {
			t.Fatal("the root node carries a fold hit")
		}
	}
}

// devJSONHasAct reports a row that folds the node at key.
func devJSONHasAct(rows []devRow, key string) bool {
	for _, row := range rows {
		if row.act == devActJSON && row.key == key {
			return true
		}
	}

	return false
}
