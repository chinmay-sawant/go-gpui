package page

import "testing"

func TestScrollRequests(t *testing.T) {
	t.Parallel()

	p := &Page{}

	p.ScrollTo(10, 20)
	req, ok := p.TakeScroll()
	if !ok || req.X != 10 || req.Y != 20 || !req.Absolute {
		t.Fatalf("scroll to %+v %v", req, ok)
	}

	if _, ok := p.TakeScroll(); ok {
		t.Fatal("scroll not cleared")
	}

	p.ScrollBy(1, 2)
	p.ScrollBy(3, 4)
	req, ok = p.TakeScroll()
	if !ok || req.X != 4 || req.Y != 6 || req.Absolute {
		t.Fatalf("scroll by %+v %v", req, ok)
	}

	p.ScrollTo(5, 6)
	p.ScrollBy(1, 1)
	req, ok = p.TakeScroll()
	if !ok || req.X != 6 || req.Y != 7 || !req.Absolute {
		t.Fatalf("scroll replace %+v %v", req, ok)
	}
}
