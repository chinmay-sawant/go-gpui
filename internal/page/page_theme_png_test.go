package page_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"
)

func TestSetThemeRejectsBrokenCSS(t *testing.T) {
	t.Parallel()

	p := newThemePage(t, "")
	if err := p.SetTheme(`#swatch{background:#ff0000`); err == nil {
		t.Fatal("unbalanced theme parsed without error")
	}
}

func TestThemeReachesPNG(t *testing.T) {
	t.Parallel()

	p := newThemePage(t, `#swatch{background:#ff0000}`)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	img, err := png.Decode(bytes.NewReader(p.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	r, g, b, _ := img.At(20, 20).RGBA()
	if r < 0xf000 || g > 0x0f00 || b > 0x0f00 {
		t.Fatalf("pixel = %x,%x,%x, want red", r, g, b)
	}
}
