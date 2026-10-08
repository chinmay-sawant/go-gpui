package window

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/blinkless/layout"
)

func TestClickWatchCounts(t *testing.T) {
	t.Parallel()

	var w clickWatch
	t0 := time.Unix(0, 0)

	if got := w.step(t0, 10, 10); got != 1 {
		t.Fatalf("first = %d", got)
	}

	if got := w.step(t0.Add(100*time.Millisecond), 11, 11); got != 2 {
		t.Fatalf("double = %d", got)
	}

	if got := w.step(t0.Add(200*time.Millisecond), 11, 12); got != 3 {
		t.Fatalf("triple = %d", got)
	}

	if got := w.step(t0.Add(300*time.Millisecond), 40, 40); got != 1 {
		t.Fatalf("far point = %d", got)
	}

	if got := w.step(t0.Add(900*time.Millisecond), 40, 40); got != 1 {
		t.Fatalf("slow click = %d", got)
	}
}

func TestEdgeScroll(t *testing.T) {
	t.Parallel()

	cases := []struct {
		y    int
		want int
	}{{-1, -48}, {0, 0}, {199, 0}, {200, 48}, {400, 48}}

	for _, c := range cases {
		if got := edgeScroll(c.y, 200, 48); got != c.want {
			t.Fatalf("edgeScroll(%d) = %d, want %d", c.y, got, c.want)
		}
	}
}

func TestDragScrollClamps(t *testing.T) {
	t.Parallel()

	app := &fakeScreen{
		width: 200, height: 200,
		boxes: []layout.Box{{X: 0, Y: 0, W: 200, H: 500}},
	}
	s := &shell{app: app, screenW: 200, screenH: 200, dragActive: true}

	s.dragScroll(200)
	if s.scrollY != 48 {
		t.Fatalf("first step = %d", s.scrollY)
	}

	for i := 0; i < 20; i++ {
		s.dragScroll(200)
	}

	if s.scrollY != 300 {
		t.Fatalf("clamped = %d", s.scrollY)
	}

	s.dragScroll(-1)
	if s.scrollY != 252 {
		t.Fatalf("back up = %d", s.scrollY)
	}

	s.dragActive = false
	s.dragScroll(200)

	if s.scrollY != 252 {
		t.Fatalf("idle drag moved = %d", s.scrollY)
	}
}
