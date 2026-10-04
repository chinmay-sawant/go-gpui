package window

import (
	"testing"
)

func TestTouchPinchSetsZoom(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 0, y: 0}, {id: 2, x: 100, y: 0}}, false)
	g.frame([]touchPos{{id: 1, x: 0, y: 0}, {id: 2, x: 150, y: 0}}, false)

	if got := g.zoomOr1(); got != 1.5 {
		t.Fatalf("zoom = %v", got)
	}

	u := g.frame(nil, false)
	if u.tap != nil {
		t.Fatal("pinch fingers tapped")
	}

	if got := g.zoomOr1(); got != 1.5 {
		t.Fatalf("zoom after release = %v", got)
	}
}

func TestPinchScale(t *testing.T) {
	t.Parallel()

	cases := []struct {
		start, now, base float64
		want             float64
	}{
		{100, 150, 1, 1.5},
		{100, 10, 1, minZoom},
		{100, 1000, 1, maxZoom},
		{0, 150, 1, 1},
		{100, 150, 2, 3},
	}

	for _, c := range cases {
		if got := pinchScale(c.start, c.now, c.base); got != c.want {
			t.Fatalf("pinchScale(%v, %v, %v) = %v", c.start, c.now, c.base, got)
		}
	}
}

func TestContentPointZoom(t *testing.T) {
	t.Parallel()

	// A box at content (50, 60) draws at screen (75, 90) under a 1.5 zoom.
	// A click there maps back to the box.
	x, y := contentPointZoom(75, 90, 0, 0, 1.5, false, 200, 200, 200, 200)
	if x != 50 || y != 60 {
		t.Fatalf("point = %v, %v", x, y)
	}

	s := &shell{
		app:     &fakeScreen{width: 200, height: 200},
		screenW: 200, screenH: 200, scrollX: 10, scrollY: 20,
	}
	s.fingers.zoom = 2

	x, y = s.contentAt(30, 40, 200, 200)
	if x != 20 || y != 30 {
		t.Fatalf("scrolled point = %v, %v", x, y)
	}
}
