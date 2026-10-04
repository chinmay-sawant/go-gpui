package window

import (
	"context"
	"testing"
)

func TestTabConsumesWhenAFieldTakesFocus(t *testing.T) {
	t.Parallel()

	app := &focusScreen{fakeScreen: &fakeScreen{}, ids: []string{"a", "b"}}
	s := &shell{app: app, ctx: context.Background()}

	taken, err := s.tabKey(true, modifiers{})
	if err != nil || !taken {
		t.Fatalf("tab = %v, %v", taken, err)
	}

	if app.focus != "a" || app.nexts != 1 {
		t.Fatalf("focus = %q nexts = %d", app.focus, app.nexts)
	}

	taken, err = s.tabKey(false, modifiers{})
	if err != nil || !taken {
		t.Fatalf("tab up = %v, %v", taken, err)
	}

	taken, err = s.tabKey(true, modifiers{Shift: true})
	if err != nil || !taken {
		t.Fatalf("shift tab = %v, %v", taken, err)
	}

	if app.focus != "b" || app.prevs != 1 {
		t.Fatalf("focus = %q prevs = %d", app.focus, app.prevs)
	}
}

func TestTabPassesThroughWithoutFields(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}

	taken, err := s.tabKey(true, modifiers{})
	if err != nil || taken {
		t.Fatalf("plain tab = %v, %v", taken, err)
	}

	app := &focusScreen{fakeScreen: &fakeScreen{}}
	s = &shell{app: app, ctx: context.Background()}

	taken, err = s.tabKey(true, modifiers{})
	if err != nil || taken {
		t.Fatalf("empty tab = %v, %v", taken, err)
	}

	if app.nexts != 1 {
		t.Fatalf("nexts = %d", app.nexts)
	}
}
