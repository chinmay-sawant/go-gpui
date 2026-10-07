package ui

import "testing"

func TestAreaCount(t *testing.T) {
	t.Parallel()

	a := Area{0, 0, 1, 1}
	if a.W() != 2 || a.H() != 2 || a.Count() != 4 || a.Empty() {
		t.Fatalf("area = %+v", a)
	}

	if !(Area{2, 0, 1, 5}).Empty() {
		t.Fatal("reversed area is not empty")
	}
}
