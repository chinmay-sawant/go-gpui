package page

import (
	"context"
	"testing"
)

func TestBindE2ECheckboxAndRadioReflectSetData(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	body := `<input type="checkbox" id="ok" data-bind="OK">` +
		`<input type="radio" id="free" name="plan" value="free" data-bind="Pick">` +
		`<input type="radio" id="pro" name="plan" value="pro" data-bind="Pick">`
	view := bdtView{}
	p := bftPage(t, body)
	bftDraw(t, p, &view)

	view.OK = true
	view.Pick = "pro"
	p.SetData(&view)
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if !p.FormChecked("ok") || !p.FormChecked("pro") || p.FormChecked("free") {
		t.Fatalf("ok=%v free=%v pro=%v", p.FormChecked("ok"), p.FormChecked("free"), p.FormChecked("pro"))
	}

	view.OK = false
	view.Pick = "free"
	p.SetData(&view)
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if p.FormChecked("ok") || !p.FormChecked("free") || p.FormChecked("pro") {
		t.Fatalf("ok=%v free=%v pro=%v", p.FormChecked("ok"), p.FormChecked("free"), p.FormChecked("pro"))
	}
}
