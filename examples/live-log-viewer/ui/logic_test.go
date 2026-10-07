package ui

import "testing"

func TestVisibleRange(t *testing.T) {
	cases := []struct {
		name                 string
		total, offset, viewH int
		start, end           int
	}{
		{"top", 100, 0, 400, 0, 35},
		{"scrolled", 100, 30 * RowH, 400, 20, 55},
		{"bottom", 100, 100*RowH - 400, 400, 65, 100},
		{"short", 5, 0, 400, 0, 5},
		{"empty", 0, 0, 400, 0, 0},
		{"negative offset", 100, -50, 400, 0, 35},
	}

	for _, c := range cases {
		start, end := visibleRange(c.total, c.offset, c.viewH, HeaderH, RowH, Overscan)
		if start != c.start || end != c.end {
			t.Errorf("%s: got %d..%d, want %d..%d", c.name, start, end, c.start, c.end)
		}
	}
}

func TestPads(t *testing.T) {
	if got := padTop(3, RowH); got != 3*RowH {
		t.Fatalf("padTop = %d", got)
	}

	if got := padBottom(10, 8, RowH); got != 2*RowH {
		t.Fatalf("padBottom = %d", got)
	}

	if got := padBottom(10, 10, RowH); got != 0 {
		t.Fatalf("padBottom end = %d", got)
	}
}

func TestPagerAnchor(t *testing.T) {
	p := Pager{}
	entries := make([]Entry, 10)
	for i := range entries {
		entries[i].ID = int64(i + 1)
	}

	p.Load(PageResult{Entries: entries})
	p.SetAnchorFromOffset(HeaderH+2*RowH+5, HeaderH)

	if p.AnchorID != 3 || p.AnchorDelta != 5 {
		t.Fatalf("anchor = %d/%d", p.AnchorID, p.AnchorDelta)
	}

	if got := p.AnchorOffset(HeaderH); got != HeaderH+2*RowH+5 {
		t.Fatalf("offset = %d", got)
	}

	p.Load(PageResult{Entries: entries[4:8]})
	if got := p.AnchorIndex(); got != 0 {
		t.Fatalf("retention anchor index = %d", got)
	}
}

func TestPagerIndexOf(t *testing.T) {
	p := Pager{Entries: []Entry{{ID: 10}, {ID: 20}, {ID: 30}}}

	if i, _ := p.IndexOf(20); i != 1 {
		t.Fatalf("exact = %d", i)
	}

	if i, _ := p.IndexOf(25); i != 2 {
		t.Fatalf("gap = %d", i)
	}

	if i, _ := p.IndexOf(99); i != 2 {
		t.Fatalf("beyond = %d", i)
	}

	if i, ok := p.IndexOf(1); i != 0 || !ok {
		t.Fatalf("before = %d", i)
	}
}
