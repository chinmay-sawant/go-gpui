package ui

import "testing"

func TestRowClickOpensDetail(t *testing.T) {
	src := &fakeSource{}
	app := newTestApp(t, src, nil)
	s := app.state

	s.table.offer(snapshotOf(fakeProc(41, "answer", 12)))
	s.table.refresh()
	clickBox(t, app, boxByAction(t, app, actProcesses))

	row := boxByAction(t, app, prefixOpen+"41:410")
	clickBox(t, app, row)

	if s.sel.id != "41:410" || !s.sel.found {
		t.Fatalf("selection = %+v", s.sel)
	}

	if len(src.tracks) == 0 || src.tracks[len(src.tracks)-1] != "41:410" {
		t.Fatalf("source tracks = %v", src.tracks)
	}
}
