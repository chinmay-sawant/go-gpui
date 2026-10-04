package page_test

import (
	"testing"
)

func TestToggleDirtiesControlAndLabel(t *testing.T) {
	t.Parallel()

	screen := formPage(t, `<label id="lab" style="display:block">Agree</label><input id="agree" type="checkbox">`)

	clickID(t, screen, "agree")

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("no dirty rect")
	}

	if !boxIn(rect, boxFor(t, screen, "agree")) || !boxIn(rect, boxFor(t, screen, "lab")) {
		t.Fatalf("rect %v misses the control or the label", rect)
	}
}
