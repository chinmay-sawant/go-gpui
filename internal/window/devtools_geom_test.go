package window

import "testing"

// TestDevScreenRectNormal checks the scroll subtraction of the one
// conversion function.
func TestDevScreenRectNormal(t *testing.T) {
	t.Parallel()

	got := devScreenRect(devRect{X: 16, Y: 32, W: 120, H: 48}, 10, 20, 1, false, 800, 600, 800, 600)
	want := devRect{X: 6, Y: 12, W: 120, H: 48}

	if got != want {
		t.Fatalf("screen rect = %+v, want %+v", got, want)
	}
}

// TestDevScreenRectZoom checks the pinch zoom multiplies before the scroll
// subtraction, matching the draw path.
func TestDevScreenRectZoom(t *testing.T) {
	t.Parallel()

	got := devScreenRect(devRect{X: 16, Y: 32, W: 120, H: 48}, 10, 20, 1.5, false, 800, 600, 800, 600)
	want := devRect{X: 14, Y: 28, W: 180, H: 72}

	if got != want {
		t.Fatalf("screen rect = %+v, want %+v", got, want)
	}
}

// TestDevScreenRectStretched checks the scale factors drawReplayScaled uses.
func TestDevScreenRectStretched(t *testing.T) {
	t.Parallel()

	got := devScreenRect(devRect{X: 16, Y: 32, W: 120, H: 48}, 0, 0, 1, true, 400, 300, 800, 600)
	want := devRect{X: 32, Y: 64, W: 240, H: 96}

	if got != want {
		t.Fatalf("screen rect = %+v, want %+v", got, want)
	}
}

// TestDevPanelNeverCoversBadge checks the panel stays clear of the fallback
// badge at the top right, on a large and on a small screen.
func TestDevPanelNeverCoversBadge(t *testing.T) {
	t.Parallel()

	for _, size := range [][2]int{{1920, 1080}, {300, 200}, {260, 120}} {
		r := devPanelRect(size[0], size[1], 260, 180, true)
		bx, by, bw, bh := badgeRect(float64(size[0]), badgeLabel)

		if r.X < bx+bw && bx < r.X+r.W && r.Y < by+bh && by < r.Y+r.H {
			t.Fatalf("panel %+v covers the badge at %v", r, size)
		}
	}
}

// TestDevPanelAboveScrollbar checks the panel bottom clears the horizontal
// scrollbar strip.
func TestDevPanelAboveScrollbar(t *testing.T) {
	t.Parallel()

	r := devPanelRect(800, 600, 260, 180, true)
	if r.Y+r.H > 600-scrollbarThickness {
		t.Fatalf("panel bottom = %v, want above the strip", r.Y+r.H)
	}
}
