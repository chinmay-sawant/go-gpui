package page_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

func TestClickRewritesThePicture(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<div id="n" data-action="add">{{.N}}</div>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	data := struct{ N int }{N: 1}
	screen.Handle(page.Handlers{
		Click: func(_ context.Context, box page.Box) error {
			if box.Action == "add" {
				data.N++
			}

			screen.SetData(data)

			return nil
		},
	})
	screen.SetData(data)

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	before := screen.Generation()
	x, y := boxCenter(t, screen, "n")
	if err := screen.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	if screen.Generation() <= before {
		t.Fatalf("generation = %d, before = %d", screen.Generation(), before)
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(screen.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width < 320 || cfg.Height < 400 {
		t.Fatalf("png = %d x %d", cfg.Width, cfg.Height)
	}

	if got := boxText(t, screen, "n"); got != "2" {
		t.Fatalf("text = %q", got)
	}
}
