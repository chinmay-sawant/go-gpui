package window

import (
	"context"
	"testing"
	"time"
)

// TestTouchTapDisarmsHold checks a touch tap ends the held-press watch. A tap
// that left the watch armed fired the long press 450 ms later, after the
// finger was already up.
func TestTouchTapDisarmsHold(t *testing.T) {
	t.Parallel()

	app := &holdScreen{selectScreen: &selectScreen{fakeScreen: &fakeScreen{}}}
	s := &shell{app: app, ctx: context.Background()}
	t0 := time.Unix(0, 0)

	s.hold.arm(t0, 10, 20)

	if err := s.tapAt(10, 20); err != nil {
		t.Fatal(err)
	}

	if err := s.holdCheck(t0.Add(longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 0 {
		t.Fatalf("holds = %d after a tap", app.holds)
	}
}
