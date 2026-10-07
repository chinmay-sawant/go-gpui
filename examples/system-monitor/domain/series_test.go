package domain

import "testing"

// TestRingWraps checks that a full ring keeps the newest points in order.
func TestRingWraps(t *testing.T) {
	r := NewRing(3)
	for i := 0; i < 5; i++ {
		r.Push(Point{Value: float64(i), Valid: true})
	}

	if r.Len() != 3 || r.Cap() != 3 {
		t.Fatalf("len=%d cap=%d", r.Len(), r.Cap())
	}

	got := r.Points()
	if len(got) != 3 || got[0].Value != 2 || got[2].Value != 4 {
		t.Fatalf("points = %+v", got)
	}
}

// TestRingReset checks that reset keeps capacity and drops points.
func TestRingReset(t *testing.T) {
	r := NewRing(2)
	r.Push(Point{Value: 1})
	r.Reset()

	if r.Len() != 0 || r.Points() != nil {
		t.Fatalf("len=%d points=%v", r.Len(), r.Points())
	}

	r.Push(Point{Value: 7})
	if got := r.Points(); len(got) != 1 || got[0].Value != 7 {
		t.Fatalf("points = %+v", got)
	}
}

// TestRingPointsCopy checks that Points returns a copy the caller may keep.
func TestRingPointsCopy(t *testing.T) {
	r := NewRing(2)
	r.Push(Point{Value: 1})

	got := r.Points()
	r.Push(Point{Value: 2})

	if got[0].Value != 1 {
		t.Fatal("Points did not return a copy")
	}
}

// TestRingZeroCapacity checks that a bad capacity still yields a usable ring.
func TestRingZeroCapacity(t *testing.T) {
	r := NewRing(0)
	r.Push(Point{Value: 1})

	if r.Len() != 1 || r.Cap() != 1 {
		t.Fatalf("len=%d cap=%d", r.Len(), r.Cap())
	}
}
