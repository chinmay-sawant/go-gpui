package ui

import "testing"

func TestComputeWindowAtTop(t *testing.T) {
	t.Parallel()

	// 1000x700 viewport: grid 948x614, so rows 0..23 and columns 0..8.
	w := computeWindow("s", 0, 0, 1000, 700, 200, 20, 2)
	if w.R0 != 0 || w.C0 != 0 {
		t.Fatalf("first = %d,%d", w.R0, w.C0)
	}

	if w.R1 != 25 || w.C1 != 10 {
		t.Fatalf("last = %d,%d, want 25,10", w.R1, w.C1)
	}
}

func TestComputeWindowScrolled(t *testing.T) {
	t.Parallel()

	// scrollY 1000 -> first visible row floor(1000/26)=38, minus overscan 2.
	w := computeWindow("s", 0, 1000, 1000, 700, 200, 20, 2)
	if w.R0 != 36 {
		t.Fatalf("R0 = %d, want 36", w.R0)
	}

	if w.R1 != 64 {
		t.Fatalf("R1 = %d, want 64", w.R1)
	}
}

func TestComputeWindowClampsToSheet(t *testing.T) {
	t.Parallel()

	w := computeWindow("s", 0, 1000000, 1000, 700, 10, 5, 4)
	if w.R0 != 5 || w.R1 != 9 || w.C0 != 0 || w.C1 != 4 {
		t.Fatalf("window = %+v", w)
	}

	w = computeWindow("s", 0, 0, 1000, 700, 0, 0, 4)
	if !w.Empty() {
		t.Fatalf("empty sheet window = %+v", w)
	}
}

func TestComputeWindowTinyView(t *testing.T) {
	t.Parallel()

	// A grid smaller than one cell still renders the cell under the offset.
	w := computeWindow("s", 0, 0, 100, 100, 200, 20, 1)
	if w.R0 != 0 || w.C0 != 0 || w.R1 < 0 || w.C1 < 0 {
		t.Fatalf("tiny window = %+v", w)
	}

	if w.R1 > 2 || w.C1 > 2 {
		t.Fatalf("tiny window too wide: %+v", w)
	}
}

func TestWindowCover(t *testing.T) {
	t.Parallel()

	w := Window{Sheet: "s", R0: 10, C0: 2, R1: 30, C1: 8}
	if !w.Cover(Area{12, 3, 20, 7}) {
		t.Fatal("cover missed an inside area")
	}

	if w.Cover(Area{9, 3, 20, 7}) || w.Cover(Area{12, 3, 31, 7}) {
		t.Fatal("cover accepted an outside area")
	}

	if (Window{R1: -1}.Cover(Area{0, 0, 0, 0})) {
		t.Fatal("empty window covered a cell")
	}
}
