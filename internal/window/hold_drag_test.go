package window

import (
	"context"
	"slices"
	"testing"
	"time"
)

func TestClaimedHoldDragsWithoutScrolling(t *testing.T) {
	t.Parallel()

	app := &holdScreen{
		selectScreen: &selectScreen{fakeScreen: &fakeScreen{width: 200, height: 200}},
		claim:        true,
	}
	s := &shell{app: app, ctx: context.Background(), screenW: 200, screenH: 200}
	t0 := time.Unix(0, 0)

	s.hold.arm(t0, 10, 10)
	if err := s.holdCheck(t0.Add(longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if !s.hold.claimed || !s.dragActive {
		t.Fatalf("claimed = %v, dragActive = %v", s.hold.claimed, s.dragActive)
	}

	s.scrollY = 40
	now := []touchPos{{id: 1, x: 10, y: 60}}

	if err := s.touchMove(touchUpdate{dy: 50}, now, 200, 200); err != nil {
		t.Fatal(err)
	}

	if s.scrollY != 40 {
		t.Fatalf("scrollY = %d, want 40", s.scrollY)
	}

	if !slices.Contains(app.calls, "drag") {
		t.Fatalf("calls = %v, want a drag", app.calls)
	}
}
