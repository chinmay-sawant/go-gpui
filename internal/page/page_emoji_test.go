package page

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestEmojiPassPaintsImages(t *testing.T) {
	ctx := context.Background()
	p, err := New(Config{HTML: `<html><body><p id="m">Good 😂 and 🎉!</p></body></html>`, Width: 400, Height: 200})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(p.source, `src="emoji/1f602"`) || !strings.Contains(p.source, `src="emoji/1f389"`) {
		t.Fatalf("no emoji images: %s", p.source)
	}

	if !strings.Contains(p.source, "data-gpui-emoji") {
		t.Fatal("no emoji sizing rule")
	}

	imgs := 0
	for i := range p.Display().Ops {
		if p.Display().Ops[i].Kind == layout.DisplayOpImage {
			imgs++
		}
	}

	if imgs != 2 {
		t.Fatalf("image ops = %d, want 2", imgs)
	}
}

func TestEmojiInFieldValue(t *testing.T) {
	ctx := context.Background()
	p, err := New(Config{HTML: `<input id="e" type="text" value="hi">`, Width: 300, Height: 100})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if err := p.Focus(ctx, "e"); err != nil {
		t.Fatal(err)
	}

	if err := p.Type(ctx, "😂"); err != nil {
		t.Fatal(err)
	}

	if got := p.FormValue("e"); got != "hi😂" {
		t.Fatalf("value = %q", got)
	}

	if !strings.Contains(p.source, `src="emoji/1f602"`) {
		t.Fatalf("field paints no emoji image: %s", p.source)
	}
}

func TestEmojiImageResolves(t *testing.T) {
	p, err := New(Config{HTML: `<p>x</p>`, Width: 100, Height: 100})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := p.imageBytes("emoji/1f602"); err != nil {
		t.Fatalf("bundled emoji: %v", err)
	}

	if _, err := p.imageBytes("emoji/1fae0"); err == nil {
		t.Fatal("unsupported emoji resolved")
	}
}
