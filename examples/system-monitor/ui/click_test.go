package ui

import (
	"context"
	"testing"
)

func TestGraphColumnsBounded(t *testing.T) {
	if got := graphColumns(10, 1); got != 24 {
		t.Fatalf("narrow graph columns = %d", got)
	}

	if got := graphColumns(100000, 1); got != graphColsMax {
		t.Fatalf("wide graph columns = %d", got)
	}
}

func TestHandleClickSortAndPaging(t *testing.T) {
	app := newTestApp(t, &fakeSource{}, nil)
	s := app.state
	s.table.offer(snapshotOf(manyProcs(120)...))
	s.table.refresh()

	ctx := context.Background()

	if err := app.handleClick(ctx, actNext); err != nil {
		t.Fatal(err)
	}

	if s.table.page != 2 {
		t.Fatalf("page after next = %d", s.table.page)
	}

	if err := app.handleClick(ctx, prefixSort+string(sortName)); err != nil {
		t.Fatal(err)
	}

	if s.table.key != sortName || s.table.page != 1 {
		t.Fatalf("sort did not reset the page: key=%q page=%d", s.table.key, s.table.page)
	}

	if rows := s.table.pageView().Rows; rows[0].Name != "proc-000" {
		t.Fatalf("name sort first row = %q", rows[0].Name)
	}
}

func TestBuildViewBoundsRows(t *testing.T) {
	app := newTestApp(t, &fakeSource{}, nil)
	s := app.state
	s.table.offer(snapshotOf(manyProcs(10000)...))
	s.table.refresh()

	v := app.buildView()
	if len(v.Rows) > rowsPerPage {
		t.Fatalf("view rendered %d rows", len(v.Rows))
	}

	if v.Total != 10000 || v.Pages != 200 || !v.HasNext || v.HasPrev {
		t.Fatalf("pager view = total=%d pages=%d prev=%v next=%v",
			v.Total, v.Pages, v.HasPrev, v.HasNext)
	}
}
