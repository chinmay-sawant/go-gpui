package window

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

type touchReleaseScreen struct {
	fakeScreen
	releases int
}

func (f *touchReleaseScreen) Release(context.Context) error { f.releases++; return nil }

func TestIdleTouchLeavesMousePressAlone(t *testing.T) {
	f := &touchReleaseScreen{fakeScreen: fakeScreen{width: 200, height: 200}}
	s := &shell{app: f, ctx: context.Background(), screenW: 200, screenH: 200, pageZoom: 1, mouseDown: true}
	if err := s.applyTouches(nil, true, 200, 200); err != nil {
		t.Fatal(err)
	}
	if f.releases != 0 {
		t.Fatal("idle touch released the mouse press")
	}
}

func TestMovingTouchReleasesOnLift(t *testing.T) {
	f := &touchReleaseScreen{fakeScreen: fakeScreen{width: 200, height: 200,
		boxes: []layout.Box{{W: 200, H: 400}}}}
	s := &shell{app: f, ctx: context.Background(), screenW: 200, screenH: 200, pageZoom: 1}
	for _, now := range [][]touchPos{{{id: 1, x: 40, y: 40}}, {{id: 1, x: 40, y: 20}}, nil} {
		if err := s.applyTouches(now, false, 200, 200); err != nil {
			t.Fatal(err)
		}
	}
	if f.releases != 2 || s.scrollY != 20 {
		t.Fatalf("releases %d scroll %d", f.releases, s.scrollY)
	}
}
