package page_test

import (
	"context"
	"testing"
)

func TestFormBlurSkipsType(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	body := `<input id="t" type="text"` + wide + `><div id="plain">x</div>`
	screen := formPage(t, body)
	clickID(t, screen, "t")
	if err := screen.Type(ctx, "ab"); err != nil {
		t.Fatal(err)
	}

	clickID(t, screen, "plain")
	if err := screen.Type(ctx, "z"); err != nil {
		t.Fatal(err)
	}

	if screen.FormValue("t") != "ab" {
		t.Fatalf("value %q", screen.FormValue("t"))
	}
}

func TestFormLoadClearsValue(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen := formPage(t, `<input id="t" type="text"`+wide+`>`)
	clickID(t, screen, "t")
	if err := screen.Type(ctx, "ab"); err != nil {
		t.Fatal(err)
	}

	if err := screen.Load(ctx, `<div id="z">Z</div>`); err != nil {
		t.Fatal(err)
	}

	if screen.FormValue("t") != "" {
		t.Fatalf("value %q", screen.FormValue("t"))
	}
}
