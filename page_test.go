package gpui_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

func TestNewRejectsEmptyHTML(t *testing.T) {
	t.Parallel()

	_, err := gpui.New(gpui.Config{Width: 320, Height: 400})
	if err != gpui.ErrEmptyHTML {
		t.Fatalf("err = %v", err)
	}
}

func TestSetSizeClamps(t *testing.T) {
	t.Parallel()

	page, err := gpui.New(gpui.Config{
		HTML:      "<p>Hi</p>",
		Width:     480,
		Height:    640,
		MinWidth:  320,
		MinHeight: 400,
		MaxWidth:  800,
		MaxHeight: 900,
	})
	if err != nil {
		t.Fatal(err)
	}

	page.SetSize(10, 10)
	width, height := page.Size()
	if width != 320 || height != 400 {
		t.Fatalf("min clamp = %d x %d", width, height)
	}

	page.SetSize(5000, 5000)
	width, height = page.Size()
	if width != 800 || height != 900 {
		t.Fatalf("max clamp = %d x %d", width, height)
	}
}

func TestImageBeforeRedrawIsNil(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	page, err := gpui.New(gpui.Config{
		HTML:   "<p>Hi</p>",
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	if page.Image() != nil {
		t.Fatal("image before redraw")
	}

	if page.PNG() != nil {
		t.Fatal("png before redraw")
	}

	if err := page.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	img := page.Image()
	if img == nil {
		t.Fatal("image is nil")
	}

	bounds := img.Bounds()
	if bounds.Dx() < 320 || bounds.Dy() < 400 {
		t.Fatalf("image = %d x %d", bounds.Dx(), bounds.Dy())
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(page.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != bounds.Dx() || cfg.Height != bounds.Dy() {
		t.Fatalf("png = %d x %d, image = %d x %d", cfg.Width, cfg.Height, bounds.Dx(), bounds.Dy())
	}
}

func TestClickRewritesThePicture(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	page, err := gpui.New(gpui.Config{
		HTML:   `<div id="n" data-action="add">{{.N}}</div>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	data := struct{ N int }{N: 1}
	page.Handle(gpui.Handlers{
		Click: func(_ context.Context, box gpui.Box) error {
			if box.Action == "add" {
				data.N++
			}

			page.SetData(data)

			return nil
		},
	})
	page.SetData(data)

	if err := page.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	before := page.Generation()
	x, y := boxCenter(t, page, "n")
	if err := page.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	if page.Generation() <= before {
		t.Fatalf("generation = %d, before = %d", page.Generation(), before)
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(page.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width < 320 || cfg.Height < 400 {
		t.Fatalf("png = %d x %d", cfg.Width, cfg.Height)
	}

	if got := boxText(t, page, "n"); got != "2" {
		t.Fatalf("text = %q", got)
	}
}

func boxCenter(t *testing.T, page *gpui.Page, id string) (float64, float64) {
	t.Helper()

	for _, box := range page.Boxes() {
		if box.ID != id || box.W <= 0 || box.H <= 0 {
			continue
		}

		return box.X + box.W/2, box.Y + box.H/2
	}

	t.Fatalf("no box id=%q", id)

	return 0, 0
}

func boxText(t *testing.T, page *gpui.Page, id string) string {
	t.Helper()

	for _, box := range page.Boxes() {
		if box.ID == id {
			return box.Text
		}
	}

	t.Fatalf("no box id=%q", id)

	return ""
}
