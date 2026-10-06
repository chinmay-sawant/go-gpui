package window

import (
	"context"
	"testing"
	"time"
)

func TestLongPressSlopCancels(t *testing.T) {
	t.Parallel()

	app := &holdScreen{selectScreen: &selectScreen{fakeScreen: &fakeScreen{}}}
	s := &shell{app: app, ctx: context.Background()}
	t0 := time.Unix(0, 0)

	s.hold.arm(t0, 10, 20)
	s.hold.move(14, 22)

	if err := s.holdCheck(t0.Add(longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 1 {
		t.Fatalf("a small move canceled the hold, holds = %d", app.holds)
	}

	s.hold.arm(t0, 10, 20)
	s.hold.move(19, 20)

	if err := s.holdCheck(t0.Add(longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 1 {
		t.Fatalf("a moved press fired, holds = %d", app.holds)
	}
}
