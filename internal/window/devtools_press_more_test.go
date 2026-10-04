package window

import "testing"

// TestDevPanelPressHits checks a press lands on the recorded hit rects and
// that empty panel space does nothing.
func TestDevPanelPressHits(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.dev.panel = devRect{X: 100, W: 200, H: 400}
	s.dev.hits = []devHit{{
		rect: devRect{X: 100, Y: 30, W: 200, H: 16},
		act:  devActTab,
		arg:  int(devTabOps),
	}}

	s.devPanelPress(120, 35)
	if s.dev.tab != devTabOps {
		t.Fatalf("tab = %v, want Ops", s.dev.tab)
	}

	s.devPanelPress(120, 300)
	if s.dev.tab != devTabOps {
		t.Fatal("an empty panel click changed the tab")
	}
}

// TestDevResizeFollowsThePointer checks the drag width and its clamps.
func TestDevResizeFollowsThePointer(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)

	s.devResize(s.screenW - 500)
	if s.dev.dockW != 500 {
		t.Fatalf("dock width = %v, want 500", s.dev.dockW)
	}

	s.devResize(s.screenW - 1)
	if s.dev.dockW != devDockMin {
		t.Fatalf("narrow drag = %v, want %v", s.dev.dockW, devDockMin)
	}
}
