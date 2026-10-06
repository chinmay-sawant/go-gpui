package ui

import "testing"

func TestTableRefreshFreezesRows(t *testing.T) {
	tb := newTable()
	tb.offer(snapshotOf(manyProcs(120)...))
	tb.refresh()

	first := tb.pageView()
	if len(first.Rows) != rowsPerPage || first.Total != 120 || first.Pages != 3 {
		t.Fatalf("page = %+v", first)
	}

	// A newer snapshot waits and must not change the displayed rows.
	tb.offer(snapshotOf(fakeProc(7, "solo", 99)))
	if !tb.hasNew() {
		t.Fatal("pending snapshot not reported")
	}

	second := tb.pageView()
	if second.Rows[0].Name != first.Rows[0].Name || second.Total != 120 {
		t.Fatal("displayed rows changed before refresh")
	}

	if !tb.refresh() {
		t.Fatal("refresh did not adopt the pending snapshot")
	}

	third := tb.pageView()
	if third.Total != 1 || third.Pages != 1 || len(third.Rows) != 1 {
		t.Fatalf("after refresh = %+v", third)
	}
}

func TestTablePagingBounds(t *testing.T) {
	tb := newTable()
	tb.offer(snapshotOf(manyProcs(105)...))
	tb.refresh()

	if tb.page != 1 {
		t.Fatalf("initial page = %d", tb.page)
	}

	tb.prev()
	if tb.page != 1 {
		t.Fatalf("prev below first page = %d", tb.page)
	}

	tb.next()
	tb.next()

	if tb.page != 3 {
		t.Fatalf("page after next next = %d", tb.page)
	}

	tb.next()
	if tb.page != 3 {
		t.Fatalf("next past last page = %d", tb.page)
	}

	if last := tb.pageView(); len(last.Rows) != 5 {
		t.Fatalf("last page rows = %d", len(last.Rows))
	}
}
