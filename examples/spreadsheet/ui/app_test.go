package ui

import (
	"testing"
)

func TestNewShowsSeededSheet(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	v := app.View()
	if len(v.Cells) == 0 {
		t.Fatal("no cells rendered")
	}

	a1, ok := cellBox(v, 0, 0)
	if !ok || a1.Text != "name" {
		t.Fatalf("A1 = %+v ok=%v", a1, ok)
	}

	b3, ok := cellBox(v, 2, 0)
	if !ok || b3.Text != "12" {
		t.Fatalf("formula cell = %+v ok=%v", b3, ok)
	}

	if v.TotalW != HeadW+20*ColW || v.TotalH != ChromeH+200*RowH {
		t.Fatalf("content size = %d,%d", v.TotalW, v.TotalH)
	}

	if v.Ref != "A1" || v.Formula != "name" {
		t.Fatalf("formula bar = %q %q", v.Ref, v.Formula)
	}
}

func TestCloseStopsWorker(t *testing.T) {
	b := newFake()
	app, err := New(Options{Backend: b, Width: 800, Height: 600})
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Close(); err != nil {
		t.Fatal(err)
	}

	if !b.closed {
		t.Fatal("backend was not closed")
	}

	if err := app.Close(); err != nil {
		t.Fatal("second close failed")
	}
}
