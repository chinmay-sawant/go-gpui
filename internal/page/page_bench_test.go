package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

// benchHTML is the fixture for the Redraw stage benchmarks: one block fill, a
// text run, and a few rules, laid out at 640x480.
const benchHTML = `<style>` +
	`#box{width:200px;height:80px;background:#3366cc}` +
	`p{color:#222222;font-size:14px;line-height:20px}` +
	`</style><div id="box"></div>` +
	`<p>Resize me. Resize me. Resize me. Resize me. Resize me.</p>`

const benchWidth, benchHeight = 640, 480

// BenchmarkRedrawWarm measures a Redraw whose data never changes: after the
// cache lands this is a relayout plus a layout, before it is the whole
// pipeline.
func BenchmarkRedrawWarm(b *testing.B) {
	ctx := context.Background()

	screen, err := page.New(page.Config{HTML: benchHTML, Width: benchWidth, Height: benchHeight})
	if err != nil {
		b.Fatal(err)
	}

	if err := screen.Redraw(ctx); err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		if err := screen.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRedrawCold measures a Redraw after SetData, which drops the cache
// and pays for the template, the parse, and the cascade again.
func BenchmarkRedrawCold(b *testing.B) {
	ctx := context.Background()
	data := struct{ N int }{}

	screen, err := page.New(page.Config{HTML: benchHTML, Width: benchWidth, Height: benchHeight})
	if err != nil {
		b.Fatal(err)
	}

	if err := screen.Redraw(ctx); err != nil {
		b.Fatal(err)
	}

	for i := 0; b.Loop(); i++ {
		data.N = i
		screen.SetData(data)

		if err := screen.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}
