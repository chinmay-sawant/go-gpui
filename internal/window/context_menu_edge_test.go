package window

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

func TestMenuOutsideClickCloses(t *testing.T) {
	t.Parallel()

	app := &menuScreen{fakeScreen: &fakeScreen{}, items: []host.MenuItem{
		{ID: "copy", Label: "Copy", Enabled: true},
	}}
	s := &shell{app: app, ctx: context.Background(), screenW: 300, screenH: 300}
	s.openMenu(10, 20)

	handled, err := s.menuClick(250, 250)
	if err != nil || !handled {
		t.Fatalf("click = %v, %v", handled, err)
	}

	if s.menu.open {
		t.Fatal("menu stayed open")
	}
}

func TestMenuEmptyRowsStayClosed(t *testing.T) {
	t.Parallel()

	app := &menuScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: app, ctx: context.Background(), screenW: 300, screenH: 300}
	s.openMenu(10, 20)

	if s.menu.open {
		t.Fatal("menu opened with no rows")
	}
}

func TestMenuPlainScreenStaysClosed(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background(), screenW: 300, screenH: 300}
	s.openMenu(10, 20)

	if s.menu.open {
		t.Fatal("menu opened without a context menu")
	}
}

func TestMenuRectPullsInsideTheScreen(t *testing.T) {
	t.Parallel()

	app := &menuScreen{fakeScreen: &fakeScreen{}, items: []host.MenuItem{
		{ID: "copy", Label: "Copy", Enabled: true},
	}}
	s := &shell{app: app, ctx: context.Background(), screenW: 120, screenH: 80}
	s.openMenu(110, 70)

	x, y, w, h := s.menuRect()
	if x+w > 120 || y+h > 80 {
		t.Fatalf("rect = %v, %v, %v, %v", x, y, w, h)
	}

	if s.menuIndex(int(x)+1, int(y)+1) != 0 {
		t.Fatal("first row missed")
	}
}
