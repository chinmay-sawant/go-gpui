package ui

import "testing"

func TestTableFilterResetsPage(t *testing.T) {
	tb := newTable()
	tb.offer(snapshotOf(manyProcs(200)...))
	tb.refresh()
	tb.next()
	tb.next()

	if tb.page != 3 {
		t.Fatalf("setup page = %d", tb.page)
	}

	tb.setQuery("proc-00")
	if tb.page != 1 {
		t.Fatalf("filter did not reset page: %d", tb.page)
	}

	pv := tb.pageView()
	if pv.Shown != 10 || pv.Pages != 1 {
		t.Fatalf("filtered view = %+v", pv)
	}

	tb.setQuery("")
	if tb.pageView().Shown != 200 {
		t.Fatal("clearing the query lost rows")
	}
}

func TestTableQueryMatchesUserAndPID(t *testing.T) {
	tb := newTable()
	tb.offer(ProcSnapshot{Procs: []Process{
		{ID: "a", PID: 4242, Name: "alpha", User: "root"},
		{ID: "b", PID: 7, Name: "beta", User: "chinmay"},
	}})
	tb.refresh()

	tb.setQuery("root")
	if pv := tb.pageView(); pv.Shown != 1 || pv.Rows[0].Name != "alpha" {
		t.Fatalf("user query = %+v", pv)
	}

	tb.setQuery("4242")
	if pv := tb.pageView(); pv.Shown != 1 || pv.Rows[0].Name != "alpha" {
		t.Fatalf("pid query = %+v", pv)
	}

	tb.setQuery("BETA")
	if pv := tb.pageView(); pv.Shown != 1 || pv.Rows[0].Name != "beta" {
		t.Fatalf("case-insensitive query = %+v", pv)
	}
}
