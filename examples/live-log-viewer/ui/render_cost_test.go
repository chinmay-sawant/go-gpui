package ui

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestRenderLongRowStaysFixedHeight(t *testing.T) {
	a := newHeadless(t)
	ctx := context.Background()

	entries := seedEntries(20)
	for i := range entries {
		entries[i].Text = strings.Repeat("wörld🚀", 80) + "\tno-space" + strings.Repeat("x", 60)
		entries[i].Lines = 4
	}

	a.pager.Load(PageResult{Entries: entries, Total: 20})
	a.Pin(0, 400)

	if err := a.draw(ctx); err != nil {
		t.Fatal(err)
	}

	for _, b := range a.page.Boxes() {
		if strings.HasPrefix(b.ID, "row-") && b.H != RowH {
			t.Fatalf("row %s height = %v", b.ID, b.H)
		}
	}
}

func TestMeasureRedrawCost(t *testing.T) {
	a := newHeadless(t)
	a.perf = true
	ctx := context.Background()
	a.pager.Load(PageResult{Entries: seedEntries(PageLimit), Total: PageLimit, HasOlder: true})
	a.Pin(0, 720)

	start := time.Now()
	for i := 0; i < 3; i++ {
		if err := a.draw(ctx); err != nil {
			t.Fatal(err)
		}
	}

	t.Logf("3 full redraws of %d rows took %v (last %v)", PageLimit, time.Since(start), a.lastDraw)
}
