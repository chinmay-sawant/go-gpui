package window

import "testing"

func TestTouchSwipeUpLifts(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 100}}, false)
	g.frame([]touchPos{{id: 1, x: 10, y: 40}}, false)

	u := g.frame(nil, false)
	if u.swipe == nil || u.swipe.dx != 0 || u.swipe.dy != -60 {
		t.Fatalf("swipe = %+v", u.swipe)
	}

	if u.tap != nil {
		t.Fatalf("tap = %+v", u.tap)
	}
}

func TestTouchSwipeDownLifts(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 40}}, false)
	g.frame([]touchPos{{id: 1, x: 10, y: 100}}, false)

	u := g.frame(nil, false)
	if u.swipe == nil || u.swipe.dx != 0 || u.swipe.dy != 60 {
		t.Fatalf("swipe = %+v", u.swipe)
	}
}

func TestTouchSwipeNeedsDistance(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 100}}, false)
	g.frame([]touchPos{{id: 1, x: 10, y: 88}}, false)

	u := g.frame(nil, false)
	if u.swipe != nil || u.tap != nil {
		t.Fatalf("swipe = %+v tap = %+v", u.swipe, u.tap)
	}
}

func TestTouchSwipeSuppressedByMulti(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 100}, {id: 2, x: 100, y: 100}}, false)
	g.frame([]touchPos{{id: 1, x: 10, y: 40}, {id: 2, x: 100, y: 40}}, false)

	u := g.frame(nil, false)
	if u.swipe != nil || u.tap != nil {
		t.Fatalf("swipe = %+v tap = %+v", u.swipe, u.tap)
	}
}

func TestTouchSwipeSuppressedByEaten(t *testing.T) {
	t.Parallel()

	var g touchGesture
	g.frame([]touchPos{{id: 1, x: 10, y: 100}}, true)
	g.frame([]touchPos{{id: 1, x: 10, y: 40}}, false)

	u := g.frame(nil, false)
	if u.swipe != nil || u.tap != nil {
		t.Fatalf("swipe = %+v tap = %+v", u.swipe, u.tap)
	}
}
