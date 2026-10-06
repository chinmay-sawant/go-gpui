package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

func TestKeyHandlersRunWithoutDrawing(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{HTML: `<p>hi</p>`, Width: 80, Height: 40})
	if err != nil {
		t.Fatal(err)
	}

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	var got []string

	screen.Handle(page.Handlers{
		KeyDown: func(_ context.Context, key string) error {
			got = append(got, "down "+key)

			return nil
		},
		KeyUp: func(_ context.Context, key string) error {
			got = append(got, "up "+key)

			return nil
		},
	})

	before := screen.Generation()
	if err := screen.KeyDown(ctx, "space"); err != nil {
		t.Fatal(err)
	}

	if err := screen.KeyUp(ctx, "space"); err != nil {
		t.Fatal(err)
	}

	if len(got) != 2 || got[0] != "down space" || got[1] != "up space" {
		t.Fatalf("handlers saw %q", got)
	}

	if screen.Generation() != before {
		t.Fatal("a key handler drew the page")
	}
}

func TestNilKeyHandlersDoNothing(t *testing.T) {
	t.Parallel()

	screen, err := page.New(page.Config{HTML: `<p>hi</p>`, Width: 80, Height: 40})
	if err != nil {
		t.Fatal(err)
	}

	if err := screen.KeyDown(context.Background(), "space"); err != nil {
		t.Fatal(err)
	}

	if err := screen.KeyUp(context.Background(), "space"); err != nil {
		t.Fatal(err)
	}
}
