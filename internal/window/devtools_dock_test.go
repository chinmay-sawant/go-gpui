package window

import "testing"

// TestDevDockOnTheRight checks the dock covers the right edge at full
// height and clamps a width wider than the window.
func TestDevDockOnTheRight(t *testing.T) {
	t.Parallel()

	if got := devDockRect(800, 600, 340); got != (devRect{X: 460, Y: 0, W: 340, H: 600}) {
		t.Fatalf("dock = %+v, want the right 340 columns", got)
	}

	if got := devDockRect(300, 200, 340); got != (devRect{X: 0, Y: 0, W: 300, H: 200}) {
		t.Fatalf("clamped dock = %+v, want the whole screen", got)
	}
}

// TestDevBadgeClearsTheDock checks the fallback badge shifts left of the
// dock instead of hiding under it.
func TestDevBadgeClearsTheDock(t *testing.T) {
	t.Parallel()

	dock := devDockRect(800, 600, devDockWidth)
	x, y, w, h := badgeRectInset(800, badgeLabel, dock.W)
	badge := devRect{X: x, Y: y, W: w, H: h}

	if badge.X+badge.W > dock.X {
		t.Fatalf("badge %+v overlaps dock %+v", badge, dock)
	}
}

// TestDevDockWidthClamps checks the stored drag width never goes under the
// minimum or wider than the window.
func TestDevDockWidthClamps(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)

	if got := s.devDockWidth(); got != devDockWidth {
		t.Fatalf("default width = %v, want %v", got, devDockWidth)
	}

	s.dev.dockW = 100
	if got := s.devDockWidth(); got != devDockMin {
		t.Fatalf("narrow width = %v, want %v", got, devDockMin)
	}

	s.dev.dockW = 9000
	if got := s.devDockWidth(); got != float64(s.screenW) {
		t.Fatalf("wide width = %v, want the screen %d", got, s.screenW)
	}
}
