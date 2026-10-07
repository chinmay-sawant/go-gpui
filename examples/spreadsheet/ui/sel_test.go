package ui

import "testing"

func TestSelectionRectAndContains(t *testing.T) {
	t.Parallel()

	s := Selection{AnchorR: 5, AnchorC: 2, ActiveR: 1, ActiveC: 7}
	r0, c0, r1, c1 := s.Rect()
	if r0 != 1 || c0 != 2 || r1 != 5 || c1 != 7 {
		t.Fatalf("rect = %d,%d,%d,%d", r0, c0, r1, c1)
	}

	if !s.Contains(3, 4) || s.Contains(0, 4) || s.Contains(3, 8) {
		t.Fatal("contains is wrong")
	}

	if got := s.Area(); got != (Area{1, 2, 5, 7}) {
		t.Fatalf("area = %+v", got)
	}
}

func TestSelectionMoveDropsAnchor(t *testing.T) {
	t.Parallel()

	s := Selection{AnchorR: 5, AnchorC: 2, ActiveR: 1, ActiveC: 7}
	m := s.Move(1, -1, 10, 10)
	if m != (Selection{2, 6, 2, 6}) {
		t.Fatalf("move = %+v", m)
	}
}

func TestSelectionExtendKeepsAnchor(t *testing.T) {
	t.Parallel()

	s := newSelection(5, 5)
	e := s.Extend(2, -1, 10, 10)
	if e.AnchorR != 5 || e.AnchorC != 5 || e.ActiveR != 7 || e.ActiveC != 4 {
		t.Fatalf("extend = %+v", e)
	}

	if !e.Contains(5, 4) || !e.Contains(7, 5) || e.Contains(7, 6) {
		t.Fatal("extended contains is wrong")
	}
}

func TestSelectionClamp(t *testing.T) {
	t.Parallel()

	s := Selection{AnchorR: 500, AnchorC: 90, ActiveR: 400, ActiveC: 80}
	c := s.Clamp(100, 20)
	if c.AnchorR != 99 || c.AnchorC != 19 || c.ActiveR != 99 || c.ActiveC != 19 {
		t.Fatalf("clamp = %+v", c)
	}
}
