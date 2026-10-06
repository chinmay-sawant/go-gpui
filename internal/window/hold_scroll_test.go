package window

import (
	"context"
	"slices"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestUnclaimedTouchMoveScrolls(t *testing.T) {
	t.Parallel()

	app := &holdScreen{selectScreen: &selectScreen{fakeScreen: &fakeScreen{
		width: 200, height: 200,
		boxes: []layout.Box{{X: 0, Y: 0, W: 200, H: 500}},
	}}}
	s := &shell{app: app, ctx: context.Background(), screenW: 200, screenH: 200}

	now := []touchPos{{id: 1, x: 10, y: 60}}

	if err := s.touchMove(touchUpdate{dy: -50}, now, 200, 200); err != nil {
		t.Fatal(err)
	}

	if s.scrollY != 50 {
		t.Fatalf("scrollY = %d, want 50", s.scrollY)
	}

	if slices.Contains(app.calls, "drag") {
		t.Fatalf("calls = %v, want no drag", app.calls)
	}
}
