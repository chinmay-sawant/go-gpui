package ui

import "testing"

func TestStaleFetchDiscarded(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	gen := app.fetchGen
	stale := result{kind: jobFetch, gen: gen - 1, sheet: "s1", area: Area{0, 0, 0, 0}, cells: []Cell{{Raw: "stale"}}}
	if app.applyResult(stale) {
		t.Fatal("stale fetch painted")
	}

	if cell, _ := app.cellAt(0, 0); cell.Raw == "stale" {
		t.Fatal("stale fetch stored")
	}

	fresh := result{kind: jobFetch, gen: gen, sheet: "s1", area: Area{0, 0, 0, 0}, cells: []Cell{{Raw: "fresh"}}}
	if !app.applyResult(fresh) {
		t.Fatal("fresh fetch did not paint")
	}

	if cell, _ := app.cellAt(0, 0); cell.Raw != "fresh" {
		t.Fatalf("cell = %+v", cell)
	}
}

func TestFetchForOtherSheetDiscarded(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	r := result{kind: jobFetch, gen: app.fetchGen, sheet: "s2", area: Area{0, 0, 0, 0}, cells: []Cell{{Raw: "wrong"}}}
	if app.applyResult(r) {
		t.Fatal("fetch for another sheet painted")
	}
}

func TestFetchMissesWindowIsPaintOnly(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	r := result{kind: jobFetch, gen: app.fetchGen, sheet: "s1", area: Area{150, 15, 150, 15}, cells: []Cell{{Raw: "far"}}}
	if app.applyResult(r) {
		t.Fatal("offscreen fetch asked for a redraw")
	}

	if cell, _ := app.cellAt(150, 15); cell.Raw != "far" {
		t.Fatal("offscreen fetch was not cached")
	}
}

func TestQueuedFetchCoalesced(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	app.work.post(job{kind: jobFetch, gen: 100, sheet: "s1", area: Area{0, 0, 0, 0}})
	app.work.post(job{kind: jobFetch, gen: 101, sheet: "s1", area: Area{1, 1, 1, 1}})

	app.work.mu.Lock()
	fetches := 0
	for _, j := range app.work.queue {
		if j.kind == jobFetch {
			fetches++
		}
	}
	app.work.mu.Unlock()

	if fetches > 1 {
		t.Fatalf("queued fetches = %d", fetches)
	}
}
