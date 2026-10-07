package ui

import "testing"

func TestPagerStaleDiscard(t *testing.T) {
	p := NewPager(2)
	first := p.Request()
	row := Row{ID: "a"}

	if ok, _ := p.Accept(PageResponse{Gen: first.Gen, Rows: []Row{row}}); !ok {
		t.Fatal("fresh page rejected")
	}

	second := p.Request()
	if ok, _ := p.Accept(PageResponse{Gen: first.Gen, Rows: nil}); ok {
		t.Fatal("stale page accepted")
	}

	if ok, _ := p.Accept(PageResponse{Gen: second.Gen, Rows: []Row{row}}); !ok {
		t.Fatal("current page rejected")
	}
}

func TestPagerFilterReset(t *testing.T) {
	p := NewPager(50)
	p.Select("keep-me")
	first := p.Request()
	p.Accept(PageResponse{Gen: first.Gen, Rows: []Row{{ID: "a"}}, Next: "c1", Total: 3})
	p.Next()
	second := p.Request()

	if !p.SetFilter(FilterFailed) {
		t.Fatal("filter change not reported")
	}

	if p.HasPrev() || p.HasNext() {
		t.Fatal("cursors not reset")
	}

	if p.Selected() != "keep-me" {
		t.Fatalf("selection lost: %q", p.Selected())
	}

	third := p.Request()
	if ok, _ := p.Accept(PageResponse{Gen: second.Gen}); ok {
		t.Fatal("old-filter page accepted")
	}

	if ok, _ := p.Accept(PageResponse{Gen: third.Gen, Rows: []Row{{ID: "f"}}}); !ok {
		t.Fatal("new-filter page rejected")
	}

	if p.SetFilter(FilterFailed) {
		t.Fatal("same filter reported a change")
	}
}
