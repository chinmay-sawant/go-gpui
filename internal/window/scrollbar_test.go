package window

import "testing"

func TestScrollbarThumb(t *testing.T) {
	t.Parallel()

	pos, length := scrollbarThumb(100, 200, 100, 0)
	if pos != 0 || length != 50 {
		t.Fatalf("start pos=%v len=%v", pos, length)
	}

	pos, _ = scrollbarThumb(100, 200, 100, 100)
	if pos != 50 {
		t.Fatalf("end pos=%v", pos)
	}

	_, length = scrollbarThumb(100, 10000, 100, 0)
	if length != scrollbarMinThumb {
		t.Fatalf("min length=%v", length)
	}

	for _, off := range []int{0, 33, 100} {
		pos, _ := scrollbarThumb(100, 200, 100, off)
		if got := scrollbarOffset(float64(pos), 100, 200, 100); got != off {
			t.Fatalf("offset %d pos=%v round trip = %d", off, pos, got)
		}
	}
}

func TestScrollbarVisible(t *testing.T) {
	t.Parallel()

	if !scrollbarVisible(200, 100) || scrollbarVisible(100, 100) || scrollbarVisible(50, 0) {
		t.Fatal("visible check")
	}
}
