package ui

import "testing"

func TestPagerWalkBack(t *testing.T) {
	p := NewPager(2)
	first := p.Request()
	p.Accept(PageResponse{Gen: first.Gen, Rows: []Row{{ID: "a"}, {ID: "b"}}, Next: "c1", Total: 2})

	if !p.Next() {
		t.Fatal("no next page")
	}

	req := p.Request()
	ok, retry := p.Accept(PageResponse{Gen: req.Gen, Rows: nil})
	if ok || !retry {
		t.Fatalf("empty last page: ok=%v retry=%v", ok, retry)
	}

	req = p.Request()
	if ok, _ := p.Accept(PageResponse{Gen: req.Gen, Rows: []Row{{ID: "a"}, {ID: "b"}}}); !ok {
		t.Fatal("walked-back page rejected")
	}

	if p.PageNumber() != 1 || p.HasPrev() {
		t.Fatalf("walked back to page %d", p.PageNumber())
	}
}

func TestPagerPages(t *testing.T) {
	p := NewPager(50)
	if p.Pages() != 1 {
		t.Fatalf("empty pages = %d", p.Pages())
	}

	req := p.Request()
	p.Accept(PageResponse{Gen: req.Gen, Rows: []Row{{ID: "a"}}, Next: "c1", Total: 101})

	if p.Pages() != 3 {
		t.Fatalf("pages = %d, want 3", p.Pages())
	}

	if !p.HasNext() || p.HasPrev() {
		t.Fatal("first-page boundaries wrong")
	}
}
