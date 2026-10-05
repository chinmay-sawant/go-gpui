package window

import (
	"context"
	"testing"
)

type stepScreen struct {
	*windowScreen
	changed bool
	calls   int
}

func (s *stepScreen) StepScrollWindow() bool { s.calls++; return s.changed }

func TestScrollReusesUnchangedWindow(t *testing.T) {
	app := &stepScreen{windowScreen: &windowScreen{fakeScreen: &fakeScreen{}, on: true}}
	s := &shell{app: app, ctx: context.Background(), scrollY: 100}
	if err := s.syncScrollWindow(); err != nil {
		t.Fatal(err)
	}
	if app.redraws != 0 || app.calls != 1 || app.oy != 100 {
		t.Fatal("unchanged window was relaid out")
	}
	app.changed = true
	s.scrollY = 200
	if err := s.syncScrollWindow(); err != nil {
		t.Fatal(err)
	}
	if app.redraws != 1 {
		t.Fatal("changed row window did not redraw")
	}
}
