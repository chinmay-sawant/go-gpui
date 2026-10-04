package page_test

import (
	"context"
	"image"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestDiffCoversWrongDeclared(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:0"><p id="count">{{.N}}</p><div id="inc" data-action="add">+</div></body>`,
		Width:  320,
		Height: 200,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetData(struct{ N int }{0})
	screen.Handle(page.Handlers{Click: func(_ context.Context, box page.Box) error {
		if box.Action == "add" {
			screen.SetData(struct{ N int }{1})
			screen.Invalidate("inc")
		}

		return nil
	}})

	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	screen.TakeDirty()

	x, y := boxCenter(t, screen, "inc")
	if err = screen.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("no dirty rect")
	}

	count := boxFor(t, screen, "count")
	center := image.Pt(int(count.X+count.W/2), int(count.Y+count.H/2))
	if !center.In(rect) {
		t.Fatalf("rect %v misses the diffed #count", rect)
	}
}
