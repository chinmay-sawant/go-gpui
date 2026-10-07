package ui

import "testing"

func TestTileCacheBounds(t *testing.T) {
	t.Parallel()

	c := newTileCache()
	c.maxTiles = 3

	for i := 0; i < 6; i++ {
		a := Area{R0: i * 10, C0: 0, R1: i*10 + 1, C1: 1}
		cells := make([]Cell, a.Count())
		c.put("s", a, cells, 0)
	}

	if len(c.tiles) > c.maxTiles {
		t.Fatalf("tiles = %d", len(c.tiles))
	}

	if c.cells > c.maxCells {
		t.Fatalf("cells = %d", c.cells)
	}
}

func TestTileCacheGetAndSet(t *testing.T) {
	t.Parallel()

	c := newTileCache()
	a := Area{R0: 5, C0: 2, R1: 6, C1: 3}
	cells := []Cell{{Raw: "a"}, {Raw: "b"}, {Raw: "c"}, {Raw: "d"}}
	c.put("s", a, cells, 7)

	cell, ok := c.get("s", 6, 3)
	if !ok || cell.Raw != "d" {
		t.Fatalf("get = %+v ok=%v", cell, ok)
	}

	if _, ok := c.get("other", 5, 2); ok {
		t.Fatal("get crossed sheets")
	}

	c.setCell("s", 5, 2, Cell{Raw: "x"})
	if cell, _ := c.get("s", 5, 2); cell.Raw != "x" {
		t.Fatalf("setCell = %+v", cell)
	}

	c.invalidate()
	if len(c.tiles) != 0 || c.cells != 0 {
		t.Fatal("invalidate kept tiles")
	}
}

func TestOverlapsWindow(t *testing.T) {
	t.Parallel()

	w := Window{Sheet: "s", R0: 10, C0: 5, R1: 20, C1: 15}
	if !overlapsWindow(Area{20, 15, 20, 15}, w) {
		t.Fatal("corner does not overlap")
	}

	if overlapsWindow(Area{21, 15, 21, 15}, w) || overlapsWindow(Area{10, 4, 10, 4}, w) {
		t.Fatal("outside area overlaps")
	}

	if overlapsWindow(Area{0, 0, 0, 0}, Window{R1: -1}) {
		t.Fatal("empty window overlaps")
	}
}
