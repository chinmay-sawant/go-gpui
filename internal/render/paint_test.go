package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

func TestPaintReturnsImageAndBox(t *testing.T) {
	t.Parallel()

	img, boxes, err := render.Paint(
		context.Background(),
		`<p id="hi" data-action="go">Hi</p>`,
		320,
		200,
	)
	if err != nil {
		t.Fatal(err)
	}

	if img == nil {
		t.Fatal("image is nil")
	}

	if img.Bounds().Empty() {
		t.Fatal("image bounds are empty")
	}

	if len(boxes) == 0 {
		t.Fatal("no boxes")
	}
}
