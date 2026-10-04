package page

import "testing"

func TestInjectThemeIntoHead(t *testing.T) {
	t.Parallel()

	source := `<html><head><title>x</title></head><body></body></html>`
	want := `<html><head><title>x</title><style>p{color:red}</style></head><body></body></html>`

	if got := injectTheme(source, `p{color:red}`); got != want {
		t.Fatalf("got %q", got)
	}
}

func TestInjectThemeWithoutHead(t *testing.T) {
	t.Parallel()

	got := injectTheme(`<p>x</p>`, `p{color:red}`)
	if got != `<style>p{color:red}</style><p>x</p>` {
		t.Fatalf("got %q", got)
	}
}

func TestInjectThemeBlank(t *testing.T) {
	t.Parallel()

	if got := injectTheme(`<p>x</p>`, "  "); got != `<p>x</p>` {
		t.Fatalf("got %q", got)
	}
}
