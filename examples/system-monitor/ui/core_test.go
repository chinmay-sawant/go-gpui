package ui

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
)

func TestCollectorFeedsOverviewAndTable(t *testing.T) {
	ctx := context.Background()
	mgr := dummyManager(t)

	app, err := New(ctx, Config{
		Source: NewCollector(mgr),
		noPump: true,
		Width:  1024,
		Height: 700,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = app.Close() })

	if err := app.page.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	waitFor(t, "first reading", mgr.HaveReading)
	app.poll(ctx)

	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if app.state.panels["cpu"].graph.Len() == 0 {
		t.Fatal("cpu graph was not fed by the collector")
	}

	if app.state.panels["cpu"].last.Text == "" || app.state.panels["cpu"].last.Text == unavailable {
		t.Fatalf("cpu reading = %q", app.state.panels["cpu"].last.Text)
	}

	waitFor(t, "first process table", func() bool {
		_, ok := mgr.Processes()

		return ok
	})

	app.poll(ctx)

	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if !app.state.table.hasNew() {
		t.Fatal("process snapshot was not offered")
	}

	app.refresh()

	pv := app.state.table.pageView()
	if pv.Total != collector.DefaultDummyProcesses || len(pv.Rows) != rowsPerPage {
		t.Fatalf("dummy table = total=%d rows=%d", pv.Total, len(pv.Rows))
	}

	first := pv.Rows[0]
	app.pick(first.ID)

	waitFor(t, "tracked detail", func() bool {
		tr := mgr.Tracked()

		return tr.ID.String() == first.ID && !tr.Loading
	})

	app.poll(ctx)

	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if !app.state.sel.trk.OK || app.state.sel.trk.ID != first.ID {
		t.Fatalf("tracked state = %+v", app.state.sel.trk)
	}
}
