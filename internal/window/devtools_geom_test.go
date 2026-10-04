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
