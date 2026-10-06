package page_test

import (
	"context"
	"image"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

func counterPage(t *testing.T, n *int) *page.Page {
	t.Helper()

	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:0"><p id="count">{{.N}}</p><div id="inc" data-action="add">+</div></body>`,
		Width:  320,
		Height: 200,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetData(struct{ N int }{*n})
	screen.Handle(page.Handlers{Click: func(_ context.Context, box page.Box) error {
		if box.Action == "add" {
			*n++
			screen.SetData(struct{ N int }{*n})
		}

		return nil
	}})

	if err := screen.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	screen.TakeDirty()

	return screen
}

func boxFor(t *testing.T, screen *page.Page, id string) page.Box {
	t.Helper()

	for _, b := range screen.Boxes() {
		if b.ID == id {
			return b
		}
	}

	t.Fatalf("no box id=%q", id)

	return page.Box{}
}

func insideBox(t *testing.T, rect image.Rectangle, b page.Box, pad int) {
	t.Helper()

	want := image.Rect(int(b.X)-pad, int(b.Y)-pad, int(b.X+b.W)+pad, int(b.Y+b.H)+pad)
	if !rect.In(want) {
		t.Fatalf("rect %v outside %v", rect, want)
	}
}

func boxIn(rect image.Rectangle, b page.Box) bool {
	return image.Rect(int(b.X), int(b.Y), int(b.X+b.W+0.5), int(b.Y+b.H+0.5)).In(rect)
}
