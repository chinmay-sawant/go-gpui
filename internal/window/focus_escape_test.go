package window

import (
	"context"
	"testing"
)

func TestEscapeClearsFocus(t *testing.T) {
	t.Parallel()

	app := &focusScreen{fakeScreen: &fakeScreen{}, focus: "a"}
	s := &shell{app: app, ctx: context.Background()}

	if err := s.escape(); err != nil {
		t.Fatal(err)
	}

	if app.focus != "" {
		t.Fatalf("focus = %q", app.focus)
	}

	if err := s.escape(); err != nil {
		t.Fatal(err)
	}

	s = &shell{app: &fakeScreen{}, ctx: context.Background()}
	if err := s.escape(); err != nil {
		t.Fatal(err)
	}
}
