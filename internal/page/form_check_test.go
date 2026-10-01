package page_test

import "testing"

func TestFormCheckbox(t *testing.T) {
	t.Parallel()

	screen := formPage(t, `<input id="c" type="checkbox">`)
	if screen.FormChecked("c") {
		t.Fatal("starts checked")
	}

	clickID(t, screen, "c")
	if !screen.FormChecked("c") {
		t.Fatal("not checked")
	}
}

func TestFormRadios(t *testing.T) {
	t.Parallel()

	body := `<input id="a" type="radio" name="g" checked style="display:block">` +
		`<input id="b" type="radio" name="g" style="display:block">`
	screen := formPage(t, body)
	if !screen.FormChecked("a") || screen.FormChecked("b") {
		t.Fatal("initial radios")
	}

	clickID(t, screen, "b")
	if screen.FormChecked("a") || !screen.FormChecked("b") {
		t.Fatal("radio group")
	}
}
