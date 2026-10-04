package window

import (
	"testing"
)

func TestTouchDragScrolls(t *testing.T) {
	t.Parallel()

	var g touchGesture
	if u := g.frame([]touchPos{{id: 1, x: 10, y: 10}}, false); u.tap != nil {
		t.Fatal("tap on a fresh finger")
	}

	u := g.frame([]touchPos{{id: 1, x: 10, y: 60}}, false)
	if u.dx != 0 || u.dy != 50 || u.tap != nil {
		t.Fatalf("drag = %d, %d tap = %v", u.dx, u.dy, u.tap)
	}

	u = g.frame(nil, false)
	if u.tap != nil {
		t.Fatal("a moved finger tapped")
	}
}

func TestTouchTapLifts(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 10}}, false)

	u := g.frame(nil, false)
	if u.tap == nil || u.tap.x != 10 || u.tap.y != 10 {
		t.Fatalf("tap = %v", u.tap)
	}
}

func TestTouchTapSuppressedByMouse(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 10}}, true)

	u := g.frame(nil, false)
	if u.tap != nil {
		t.Fatal("a suppressed finger tapped")
	}
}

func TestTouchDragClampsAtBottom(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 0, y: 200}}, false)

	u := g.frame([]touchPos{{id: 1, x: 0, y: -200}}, false)

	sx, sy := clampScroll(0-u.dx, 0-u.dy, 200, 500, 200, 200)
	if sx != 0 || sy != 300 {
		t.Fatalf("offset = %d, %d", sx, sy)
	}
}

func TestTouchSlopKeepsATap(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 10}}, false)

	u := g.frame([]touchPos{{id: 1, x: 14, y: 12}}, false)
	if u.dx != 0 || u.dy != 0 {
		t.Fatalf("slop moved = %d, %d", u.dx, u.dy)
	}

	u = g.frame(nil, false)
	if u.tap == nil {
		t.Fatal("a small move lost the tap")
	}
}
