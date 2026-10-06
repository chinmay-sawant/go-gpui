package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

func TestReloadDoesNotFireChange(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<input id="note" type="text">`)
	p := newFilePage(t, page.Config{File: path, Width: 320, Height: 200})

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	changes := 0
	p.Handle(page.Handlers{Change: func(context.Context, page.Box) error {
		changes++

		return nil
	}})

	rewriteSource(t, path, `<input id="note" type="text"><p id="extra">extra</p>`)

	if _, err := p.PollReload(ctx); err != nil {
		t.Fatal(err)
	}

	if changes != 0 {
		t.Fatalf("changes = %d, want none for a reload", changes)
	}
}
