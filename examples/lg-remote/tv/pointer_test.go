package tv

import "testing"

func TestCutApp(t *testing.T) {
	button, id, ok := cutApp("NETFLIX|netflix")
	if !ok || button != "NETFLIX" || id != "netflix" {
		t.Fatalf("%s %s %v", button, id, ok)
	}
}
