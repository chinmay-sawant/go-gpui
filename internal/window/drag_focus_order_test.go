package window

import (
	"context"
	"testing"
)

// TestPressLeavesHandlerFocus pins the order: the caret lands before the
// click handler runs, so the handler's focus and selection survive.
func TestPressLeavesHandlerFocus(t *testing.T) {
	t.Parallel()

	app := &selectScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: app, ctx: context.Background()}

	if err := s.pressAt(10, 10); err != nil {
		t.Fatal(err)
	}

	if !app.focus {
		t.Fatal("SelectAt cleared the focus the click handler set")
	}
}
