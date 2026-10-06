package window

import (
	"context"
	"testing"
	"time"
)

// holdScreen records the long-press calls the window makes. claim is the
// answer its LongPress returns.
type holdScreen struct {
	*selectScreen
	claim bool
	holds int
	x, y  float64
}

func (f *holdScreen) LongPress(_ context.Context, x, y float64) (bool, error) {
	f.holds++
	f.x, f.y = x, y

	return f.claim, nil
}

func TestLongPressFiresOnce(t *testing.T) {
	t.Parallel()

	app := &holdScreen{selectScreen: &selectScreen{fakeScreen: &fakeScreen{}}}
	s := &shell{app: app, ctx: context.Background()}
	t0 := time.Unix(0, 0)

	s.hold.arm(t0, 10, 20)

	if err := s.holdCheck(t0.Add(longPressDelay - time.Millisecond)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 0 {
		t.Fatalf("holds = %d before the delay", app.holds)
	}

	if err := s.holdCheck(t0.Add(longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 1 || app.x != 10 || app.y != 20 {
		t.Fatalf("holds = %d at %v, %v", app.holds, app.x, app.y)
	}

	if err := s.holdCheck(t0.Add(2 * longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 1 {
		t.Fatalf("holds = %d after a repeat", app.holds)
	}

	s.hold.arm(t0, 10, 20)
	s.hold.release()

	if err := s.holdCheck(t0.Add(longPressDelay)); err != nil {
		t.Fatal(err)
	}

	if app.holds != 1 {
		t.Fatalf("holds = %d after a release", app.holds)
	}
}
