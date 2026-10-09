package window

import (
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

type fittedScreen struct{ *requestScreen }

func (*fittedScreen) ViewLocked() bool { return true }

func TestLockedViewMapsFittedControlsAndRejectsScrolling(t *testing.T) {
	app := &fittedScreen{&requestScreen{fakeScreen: &fakeScreen{
		width: 100, height: 100,
		boxes: []layout.Box{{W: 200, H: 400}},
	}}}
	s := &shell{app: app, screenW: 100, screenH: 100, scrollY: 50}
	x, y := s.contentAt(50, 75, 200, 400)
	if x != 100 || y != 300 {
		t.Fatalf("fitted pointer = %.1f, %.1f", x, y)
	}
	if s.allowPageScroll() {
		t.Fatal("locked view allows scrolling")
	}
	s.pullScroll()
	if s.scrollY != 0 {
		t.Fatal("locked view retained its scroll offset")
	}
}
