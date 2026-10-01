package window

import "testing"

func TestFramePointScalesOntoThePicture(t *testing.T) {
	t.Parallel()

	x, y := framePoint(50, 25, 200, 100, 100, 50)
	if x != 100 || y != 50 {
		t.Fatalf("point = %v, %v", x, y)
	}
}

func TestFramePointMatchesWhenSizesMatch(t *testing.T) {
	t.Parallel()

	x, y := framePoint(12, 34, 480, 640, 480, 640)
	if x != 12 || y != 34 {
		t.Fatalf("point = %v, %v", x, y)
	}
}

func TestFramePointRejectsAnEmptyScreen(t *testing.T) {
	t.Parallel()

	x, y := framePoint(12, 34, 480, 640, 0, 640)
	if x != 0 || y != 0 {
		t.Fatalf("point = %v, %v", x, y)
	}
}
