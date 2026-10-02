package window

import "testing"

func TestBadgeRectSitsInTheTopRight(t *testing.T) {
	t.Parallel()

	const screenW = 640.0
	x, y, width, height := badgeRect(screenW, badgeLabel)

	if width <= 0 || height <= 0 {
		t.Fatalf("badge size = %v x %v, want positive", width, height)
	}

	if got := x + width + 8; got != screenW {
		t.Fatalf("badge right edge = %v, want %v", got, screenW)
	}

	if y != 8 {
		t.Fatalf("badge top = %v, want 8", y)
	}
}
