package ownframe

import (
	"bytes"
	"context"
	"testing"
)

// repaintCase is one page whose incremental path must match a full Redraw.
type repaintCase struct {
	name     string
	newPage  func(t *testing.T) *Page
	interact func(t *testing.T, p *Page, ctx context.Context)
}

// TestRepaintMatchesFullRedraw drives one interaction on two pages and
// compares the incremental PNG with the PNG after a full Redraw. This is the
// guard that keeps a partial repaint from being a subtly wrong picture.
func TestRepaintMatchesFullRedraw(t *testing.T) {
	t.Parallel()

	for _, c := range repaintCases() {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			inc := c.newPage(t)
			full := c.newPage(t)

			before := inc.PNG()
			c.interact(t, inc, ctx)
			c.interact(t, full, ctx)
			if err := full.Redraw(ctx); err != nil {
				t.Fatal(err)
			}

			after := inc.PNG()
			if before == nil || after == nil {
				t.Fatal("no PNG to compare")
			}

			if bytes.Equal(before, after) {
				t.Fatal("the interaction did not change the picture")
			}

			if !bytes.Equal(after, full.PNG()) {
				t.Fatal("incremental PNG differs from a full Redraw")
			}
		})
	}
}

// newPage builds a page, lets the case set its data and handlers, and lays
// it out once.
func newPage(t *testing.T, cfg Config, setup func(p *Page)) *Page {
	t.Helper()

	p, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if setup != nil {
		setup(p)
	}

	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return p
}

func clickID(t *testing.T, p *Page, ctx context.Context, id string) {
	t.Helper()

	x, y := boxCenter(t, p, id)
	if err := p.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}

func boxCenter(t *testing.T, p *Page, id string) (float64, float64) {
	t.Helper()

	for _, b := range p.Boxes() {
		if b.ID == id {
			return b.X + b.W/2, b.Y + b.H/2
		}
	}

	t.Fatalf("no box %q", id)

	return 0, 0
}
