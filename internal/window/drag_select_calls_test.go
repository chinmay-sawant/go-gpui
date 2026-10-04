package window

import (
	"context"
	"slices"
	"testing"
)

func TestPressDragRelease(t *testing.T) {
	t.Parallel()

	app := &selectScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: app, ctx: context.Background()}

	if err := s.pressAt(10, 10); err != nil {
		t.Fatal(err)
	}

	if !s.dragActive {
		t.Fatal("drag not active after a press")
	}

	if err := s.dragAt(10, 10); err != nil {
		t.Fatal(err)
	}

	if err := s.dragAt(20, 30); err != nil {
		t.Fatal(err)
	}

	if err := s.releaseAt(); err != nil {
		t.Fatal(err)
	}

	if s.dragActive {
		t.Fatal("drag active after the release")
	}

	want := []string{"press", "at", "click", "drag", "release"}
	if !slices.Equal(app.calls, want) {
		t.Fatalf("calls = %v, want %v", app.calls, want)
	}
}

func TestDoubleAndTripleClick(t *testing.T) {
	t.Parallel()

	app := &selectScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: app, ctx: context.Background()}

	for i := 0; i < 3; i++ {
		if err := s.pressAt(10+float64(i), 10); err != nil {
			t.Fatal(err)
		}
	}

	want := []string{
		"press", "at", "click",
		"press", "at", "click", "word",
		"press", "at", "click", "line",
	}

	if !slices.Equal(app.calls, want) {
		t.Fatalf("calls = %v, want %v", app.calls, want)
	}
}

func TestPlainScreenKeepsClicks(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}

	if err := s.pressAt(1, 2); err != nil {
		t.Fatal(err)
	}

	if s.dragActive {
		t.Fatal("drag active on a plain screen")
	}

	if err := s.dragAt(3, 4); err != nil {
		t.Fatal(err)
	}

	if err := s.releaseAt(); err != nil {
		t.Fatal(err)
	}
}
