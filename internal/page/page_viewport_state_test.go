package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
	"github.com/chinmay-sawant/ownframe/internal/page"
)

// hasFill reports a fill operation of that color in the display list.
func hasFill(p *page.Page, r, g, b float64) bool {
	d := p.Display()
	if d == nil {
		return false
	}

	for i := range d.Ops {
		op := &d.Ops[i]
		if op.Kind == layout.DisplayOpFillRect && op.R == r && op.G == g && op.B == b {
			return true
		}
	}

	return false
}

func TestFocusAndCheckedSurviveARelayout(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	source := `<html><head><style>body{margin:0}` +
		`#name{width:120px;height:20px}` +
		`#name:focus{background:#ff0000}` +
		`#name:focus-visible{background:#0000ff}` +
		`#ok{width:30px;height:20px}` +
		`#ok:checked{background:#00ff00}` +
		`</style></head><body>` +
		`<input id="name"><input id="ok" type="checkbox">` +
		`</body></html>`

	p := newViewportPage(t, source, 320, 200)

	x, y := boxCenter(t, p, "name")
	if err := p.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	if p.FocusedField() != "name" {
		t.Fatalf("focus = %q, want name", p.FocusedField())
	}

	p.SetFormChecked("ok", true)
	p.SetSize(640, 400)

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if !hasFill(p, 0, 0, 1) {
		t.Fatal("focus-visible style lost after the relayout")
	}

	if !hasFill(p, 0, 1, 0) {
		t.Fatal("checked style lost after the relayout")
	}
}
