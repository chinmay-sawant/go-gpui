package ui

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
)

func TestRenderDetailMode(t *testing.T) {
	a := newHeadless(t)
	ctx := context.Background()

	a.detail = Detail{Entry: seedEntries(5)[4], Lines: splitLines("first\nsecond")}
	a.detailOpen = true
	a.Pin(0, 400)

	if err := a.draw(ctx); err != nil {
		t.Fatal(err)
	}

	boxes := a.page.Boxes()
	if _, ok := boxByID(boxes, "detail-back"); !ok {
		t.Fatal("detail back button missing")
	}

	if _, ok := boxByID(boxes, "row-5"); ok {
		t.Fatal("list row rendered in detail mode")
	}
}

func TestClickSelectsRow(t *testing.T) {
	a := newHeadless(t)
	ctx := context.Background()

	if err := a.onClick(ctx, ownframe.Box{Action: "row-7"}); err != nil {
		t.Fatal(err)
	}

	if a.selected != 7 {
		t.Fatalf("selected = %d", a.selected)
	}
}
