package window

import (
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

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

func TestScrollbarOffsetClamps(t *testing.T) {
	t.Parallel()

	if got := scrollbarOffset(-50, 100, 200, 100); got != 0 {
		t.Fatalf("negative offset = %d", got)
	}

	if got := scrollbarOffset(1000, 100, 200, 100); got != 100 {
		t.Fatalf("past-end offset = %d", got)
	}
}

func TestScrollbarThumbNearlyFull(t *testing.T) {
	t.Parallel()

	// Content barely taller than the viewport: the thumb fills the track.
	pos, length := scrollbarThumb(100, 101, 100, 1)
	if length < 99 || length > 100 {
		t.Fatalf("length = %v", length)
	}

	if pos > 1.5 {
		t.Fatalf("pos = %v", pos)
	}
}

func TestContentSizeIncludesOverflow(t *testing.T) {
	t.Parallel()

	s := &shell{
		app: &fakeScreen{
			width: 320, height: 400,
			boxes: []layout.Box{{X: 0, Y: 0, W: 500, H: 700}},
		},
		screenW: 320, screenH: 400,
	}

	w, h := s.contentSize()
	if w != 500 || h != 700 {
		t.Fatalf("content = %dx%d", w, h)
	}
}
