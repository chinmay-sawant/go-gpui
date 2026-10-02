package page_test

import (
	"context"
	"strings"
	"testing"
)

func TestFormSelectCycles(t *testing.T) {
	t.Parallel()

	body := `<select id="s"` + wide + `><option>Red</option>` +
		`<option selected>Blue</option></select>`
	screen := formPage(t, body)
	clickID(t, screen, "s")
	if strings.TrimSpace(boxText(t, screen, "s")) != "Red" {
		t.Fatalf("text %q", boxText(t, screen, "s"))
	}
}

func TestFormFile(t *testing.T) {
	t.Parallel()

	screen := formPage(t, `<input id="f" type="file"`+wide+`>`)
	clickID(t, screen, "f")
	if err := screen.Type(context.Background(), "notes.txt"); err != nil {
		t.Fatal(err)
	}

	if boxText(t, screen, "f") != "notes.txt" {
		t.Fatalf("text %q", boxText(t, screen, "f"))
	}
}

func TestFormTextareaBackspace(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen := formPage(t, `<textarea id="a"`+wide+`></textarea>`)
	clickID(t, screen, "a")
	if err := screen.Type(ctx, "hi"); err != nil {
		t.Fatal(err)
	}

	if err := screen.Backspace(ctx); err != nil {
		t.Fatal(err)
	}

	if screen.FormValue("a") != "h" {
		t.Fatalf("value %q", screen.FormValue("a"))
	}
}

func TestFormDisabledCheckbox(t *testing.T) {
	t.Parallel()

	screen := formPage(t, `<input id="c" type="checkbox" disabled>`)
	clickID(t, screen, "c")
	if screen.FormChecked("c") {
		t.Fatal("disabled checked")
	}
}
