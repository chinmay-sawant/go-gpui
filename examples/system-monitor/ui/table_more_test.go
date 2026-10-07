package ui

import "testing"

func TestTableShrinkingLastPageClamps(t *testing.T) {
	tb := newTable()
	tb.offer(snapshotOf(manyProcs(120)...))
	tb.refresh()
	tb.next()
	tb.next()

	if tb.page != 3 {
		t.Fatalf("setup page = %d", tb.page)
	}

	tb.offer(snapshotOf(manyProcs(60)...))
	tb.refresh()

	if tb.page != 2 {
		t.Fatalf("page not clamped after shrink: %d", tb.page)
	}

	if pv := tb.pageView(); len(pv.Rows) != 10 || pv.Pages != 2 {
		t.Fatalf("clamped view = %+v", pv)
	}
}

func TestTableSortTieBreaksByIdentity(t *testing.T) {
	tb := newTable()
	tb.offer(snapshotOf(
		Process{ID: "b", PID: 2, Name: "two", CPU: 5, CPUKnown: true},
		Process{ID: "a", PID: 1, Name: "one", CPU: 5, CPUKnown: true},
	))
	tb.refresh()

	rows := tb.pageView().Rows
	if rows[0].ID != "a" || rows[1].ID != "b" {
		t.Fatalf("tie order = %s, %s", rows[0].ID, rows[1].ID)
	}

	tb.sortBy(sortCPU) // same key flips direction, identity stays ascending
	rows = tb.pageView().Rows

	if rows[0].ID != "a" || rows[1].ID != "b" {
		t.Fatalf("desc tie order = %s, %s", rows[0].ID, rows[1].ID)
	}
}

func TestTableRenderedRowsBounded(t *testing.T) {
	tb := newTable()
	tb.offer(snapshotOf(manyProcs(10000)...))
	tb.refresh()

	for page := 1; page <= tb.pages(); page++ {
		if n := len(tb.pageView().Rows); n > rowsPerPage {
			t.Fatalf("page %d rendered %d rows", page, n)
		}
	}
}
