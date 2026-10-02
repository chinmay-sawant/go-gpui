package page

import "testing"

func TestBindEditTogglesWriteStruct(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name, body, id string
		check          func(bedView, *Page) bool
	}{
		{"checkbox", `<input id="c" type="checkbox" data-bind="Agree">`, "c",
			func(v bedView, p *Page) bool { return v.Agree && p.FormChecked("c") }},
		{"radio", `<input id="a" type="radio" name="g" value="free" data-bind="Plan">` +
			`<input id="b" type="radio" name="g" value="pro" data-bind="Plan"` + bedWide + `>`, "b",
			func(v bedView, p *Page) bool { return v.Plan == "pro" && p.FormChecked("b") }},
		{"select", `<select id="s" data-bind="Plan"` + bedWide + `><option value="free">Free</option>` +
			`<option value="pro" selected>Pro</option></select>`, "s",
			func(v bedView, p *Page) bool { return v.Plan == "free" && p.FormValue("s") == "free" }},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			view := &bedView{}
			p := bedPage(t, tc.body, view)
			bedClick(t, p, tc.id)
			if !tc.check(*view, p) {
				t.Fatalf("view %+v", *view)
			}
		})
	}
}
