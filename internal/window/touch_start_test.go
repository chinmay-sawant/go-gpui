package window

import "testing"

func TestTouchUpdateReportsAFreshFinger(t *testing.T) {
	t.Parallel()

	var g touchGesture

	u := g.frame([]touchPos{{id: 7, x: 10, y: 20}}, false)
	if u.start == nil || u.start.id != 7 || u.start.x != 10 || u.start.y != 20 {
		t.Fatalf("start = %v", u.start)
	}

	u = g.frame([]touchPos{{id: 7, x: 30, y: 20}}, false)
	if u.start != nil {
		t.Fatalf("a moving finger restarted the hold: %v", u.start)
	}
}
