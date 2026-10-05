package window

import (
	"context"
	"slices"
	"testing"
)

// TestTapAtPlacesCaret checks a tap lands the caret before the click
// handler runs, matching pressAt.
func TestTapAtPlacesCaret(t *testing.T) {
	t.Parallel()

	app := &selectScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: app, ctx: context.Background()}

	if err := s.tapAt(10, 10); err != nil {
		t.Fatal(err)
	}

	want := []string{"press", "at", "click", "release"}
	if !slices.Equal(app.calls, want) {
		t.Fatalf("calls = %v, want %v", app.calls, want)
	}
}

// TestTapAtPlainScreen checks a page without the selector keeps the old
// press, click, release tap.
func TestTapAtPlainScreen(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}

	if err := s.tapAt(1, 2); err != nil {
		t.Fatal(err)
	}
}

// TestDevPickAltPlacesCaret checks Alt+click lands the caret before the
// click handler runs.
func TestDevPickAltPlacesCaret(t *testing.T) {
	t.Parallel()

	app := &selectScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: app, ctx: context.Background()}

	if err := s.devPick(10, 10, true); err != nil {
		t.Fatal(err)
	}

	want := []string{"press", "at", "click"}
	if !slices.Equal(app.calls, want) {
		t.Fatalf("calls = %v, want %v", app.calls, want)
	}

	if !s.dev.forward {
		t.Fatal("release not forwarded after an alt pick")
	}
}

// TestDevPickAltPlainScreen checks Alt+click still forwards on a page
// without the selector.
func TestDevPickAltPlainScreen(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}

	if err := s.devPick(1, 2, true); err != nil {
		t.Fatal(err)
	}

	if !s.dev.forward {
		t.Fatal("release not forwarded after an alt pick")
	}
}
