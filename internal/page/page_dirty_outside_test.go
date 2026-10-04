package page_test

import (
	"context"
	"image"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestDirtyRectOutsideIsFrame(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:0"><p id="t" style="margin-top:300px">{{.V}}</p></body>`,
		Width:  320,
		Height: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetData(struct{ V string }{V: "a"})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	screen.TakeDirty()

	screen.SetData(struct{ V string }{V: "b"})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok || rect != image.Rect(0, 0, 320, 100) {
		t.Fatalf("rect %v ok %v", rect, ok)
	}
}
