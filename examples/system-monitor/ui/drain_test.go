package ui

import (
	"context"
	"testing"
)

func TestDrainDiscardsStaleGeneration(t *testing.T) {
	app := newTestApp(t, &fakeSource{}, nil)
	app.mail.putSummary(app.currentGen()+1, sample("cpu", 0.5))

	if err := app.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	if app.state.panels["cpu"].graph.Len() != 0 {
		t.Fatal("a stale-generation sample was applied")
	}
}

func TestDetailStaleIdentityDiscarded(t *testing.T) {
	app := newTestApp(t, &fakeSource{}, nil)
	app.pick("a")

	app.mail.putTracked(app.currentGen(), Tracked{ID: "b", OK: true})

	if err := app.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	if app.state.sel.trk.OK {
		t.Fatal("detail for another identity was applied")
	}

	app.mail.putTracked(app.currentGen(), Tracked{ID: "a", OK: true})

	if err := app.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	if !app.state.sel.trk.OK {
		t.Fatal("detail for the selected identity was dropped")
	}
}

func TestAppCloseIsIdempotent(t *testing.T) {
	src := &fakeSource{}
	app := newTestApp(t, src, nil)

	if err := app.Close(); err != nil {
		t.Fatal(err)
	}

	if err := app.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestPollDeduplicatesSamples(t *testing.T) {
	src := &fakeSource{haveSum: true, summary: sample("cpu", 0.5)}
	app := newTestApp(t, src, nil)

	app.poll(context.Background())
	app.poll(context.Background())

	b := app.mail.take()
	if b.summary == nil {
		t.Fatal("first poll did not post the sample")
	}

	app.poll(context.Background())

	if b := app.mail.take(); b.summary != nil {
		t.Fatal("unchanged sample was reposted")
	}
}
