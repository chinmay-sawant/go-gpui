package page_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

// heavyHTML builds a page with n CSS rules so the sheet collection is a
// visible share of the Redraw cost.
func heavyHTML(n int) string {
	var b strings.Builder
	b.WriteString("<style>")

	for i := 0; i < n; i++ {
		fmt.Fprintf(&b, ".r%d{color:#%06x;font-size:%dpx;margin:%dpx}", i, i*7919&0xFFFFFF, 10+i%8, i%5)
	}

	b.WriteString("</style>")
	b.WriteString(`<p class="r7">heavy stylesheet</p><div id="box"></div>`)

	return b.String()
}

// BenchmarkRedrawHeavyWarm measures repeated Redraws of a page with 200 CSS
// rules and unchanged data.
func BenchmarkRedrawHeavyWarm(b *testing.B) {
	ctx := context.Background()

	screen, err := page.New(page.Config{HTML: heavyHTML(200), Width: benchWidth, Height: benchHeight})
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
