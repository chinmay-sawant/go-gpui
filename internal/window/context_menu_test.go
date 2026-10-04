package window

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/clipboard"
	"github.com/chinmay-sawant/go-gpui/internal/host"
)

func TestMenuPasteFromRow(t *testing.T) {
	clipboard.UseMemory(true)
	clipboard.Write("hello\n")

	app := &menuScreen{fakeScreen: &fakeScreen{width: 300, height: 300}, items: []host.MenuItem{
		{ID: "cut", Label: "Cut", Enabled: false},
		{ID: "paste", Label: "Paste", Enabled: true},
	}}
	s := &shell{app: app, ctx: context.Background(), screenW: 300, screenH: 300}
	s.openMenu(10, 20)

	if !s.menu.open || len(s.menu.items) != 2 {
		t.Fatalf("menu = %+v", s.menu)
	}

	row := s.menuIndex(15, 20+menuRowH+5)
	if row != 1 {
		t.Fatalf("row = %d", row)
	}

	handled, err := s.menuClick(15, 20+menuRowH+5)
	if err != nil || !handled {
		t.Fatalf("click = %v, %v", handled, err)
	}

	if app.pasted != "hello" {
		t.Fatalf("pasted = %q", app.pasted)
	}

	if s.menu.open {
		t.Fatal("menu stayed open")
	}
}

func TestMenuSkipsDisabledRow(t *testing.T) {
	t.Parallel()

	app := &menuScreen{fakeScreen: &fakeScreen{}, items: []host.MenuItem{
		{ID: "undo", Label: "Undo", Enabled: false},
	}}
	s := &shell{app: app, ctx: context.Background(), screenW: 300, screenH: 300}
	s.openMenu(10, 20)

	handled, err := s.menuClick(15, 25)
	if err != nil || !handled {
		t.Fatalf("click = %v, %v", handled, err)
	}

	if app.pasted != "" || s.menu.open {
		t.Fatal("disabled row acted or kept the menu")
	}
}

func TestEscapeClosesMenu(t *testing.T) {
	t.Parallel()

	app := &menuScreen{fakeScreen: &fakeScreen{}, items: []host.MenuItem{
		{ID: "copy", Label: "Copy", Enabled: true},
	}}
	s := &shell{app: app, ctx: context.Background(), screenW: 300, screenH: 300}
	s.openMenu(5, 5)

	if err := s.escape(); err != nil {
		t.Fatal(err)
	}

	if s.menu.open {
		t.Fatal("menu stayed open after escape")
	}
}
