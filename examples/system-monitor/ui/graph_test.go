package ui

import "testing"

func TestGraphRingBounded(t *testing.T) {
	g := NewGraph(3)

	for i := 1; i <= 5; i++ {
		g.Push(float64(i), true)
	}

	if g.Len() != 3 {
		t.Fatalf("Len = %d, want 3", g.Len())
	}

	vals, oks := g.Columns(3)

	want := []float64{3, 4, 5}

	for i, v := range vals {
		if !oks[i] || v != want[i] {
			t.Errorf("column %d = %v ok=%v, want %v", i, v, oks[i], want[i])
		}
	}

	if v, ok := g.Last(); !ok || v != 5 {
		t.Errorf("Last = %v ok=%v", v, ok)
	}

	if v, ok := g.Max(); !ok || v != 5 {
		t.Errorf("Max = %v ok=%v", v, ok)
	}
}

func TestGraphReset(t *testing.T) {
	g := NewGraph(4)
	g.Push(4, true)
	g.Reset()

	if g.Len() != 0 {
		t.Fatalf("Len after reset = %d", g.Len())
	}

	if _, ok := g.Last(); ok {
		t.Fatal("Last after reset is known")
	}
}

func TestGraphColumnsPreservePeak(t *testing.T) {
	g := NewGraph(6)

	for _, v := range []float64{1, 2, 9, 3, 4, 5} {
		g.Push(v, true)
	}

	vals, oks := g.Columns(3)
	want := []float64{2, 9, 5}

	for i, v := range vals {
		if !oks[i] || v != want[i] {
			t.Errorf("column %d = %v ok=%v, want %v", i, v, oks[i], want[i])
		}
	}
}

func TestGraphColumnsFewerSamples(t *testing.T) {
	g := NewGraph(8)
	g.Push(7, true)
	g.Push(8, true)

	vals, oks := g.Columns(4)
	if oks[0] || oks[1] {
		t.Fatalf("leading columns should be gaps: %v", oks)
	}

	if !oks[2] || vals[2] != 7 || !oks[3] || vals[3] != 8 {
		t.Fatalf("right-aligned columns = %v %v", vals, oks)
	}
}

func TestGraphColumnsGap(t *testing.T) {
	g := NewGraph(2)
	g.Push(1, true)
	g.Push(0, false)

	vals, oks := g.Columns(2)
	if !oks[0] || vals[0] != 1 {
		t.Errorf("column 0 = %v ok=%v", vals[0], oks[0])
	}

	if oks[1] {
		t.Errorf("gap column = %v", vals[1])
	}
}
