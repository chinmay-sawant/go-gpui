package ui

import (
	"context"
	"testing"
)

func TestRefreshKeepsSelectionByIdentity(t *testing.T) {
	app := newTestApp(t, &fakeSource{}, nil)
	s := app.state

	s.table.offer(snapshotOf(fakeProc(10, "ten", 5)))
	s.table.refresh()
	app.pick("10:100")

	s.table.offer(snapshotOf(fakeProc(10, "ten", 42)))
	app.refresh()

	if !s.sel.found || s.sel.proc.CPU != 42 {
		t.Fatalf("selection not updated by identity: %+v", s.sel)
	}

	s.table.offer(snapshotOf(fakeProc(11, "eleven", 1)))
	app.refresh()

	if s.sel.found || s.sel.proc.Name != "ten" {
		t.Fatalf("exited selection lost its last-known row: %+v", s.sel)
	}
}

func TestToggleModeResetsBaselines(t *testing.T) {
	src := &fakeSource{}
	app := newTestApp(t, src, nil)
	s := app.state

	app.mail.putSummary(app.currentGen(), sample("cpu", 0.5))

	if err := app.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	if s.panels["cpu"].graph.Len() != 1 {
		t.Fatal("setup: sample not applied")
	}

	before := app.currentGen()

	if err := app.handleClick(context.Background(), actMode); err != nil {
		t.Fatal(err)
	}

	if !src.liveMode() {
		t.Fatal("source did not switch to live")
	}

	if s.panels["cpu"].graph.Len() != 0 || s.haveAny {
		t.Fatal("graphs were not reset on the mode switch")
	}

	if app.currentGen() == before {
		t.Fatal("generation did not advance on the mode switch")
	}

	if err := app.handleClick(context.Background(), actMode); err != nil {
		t.Fatal(err)
	}

	if src.liveMode() {
		t.Fatal("source did not switch back to dummy")
	}
}
