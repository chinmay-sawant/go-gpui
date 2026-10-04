package page_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// writeSource writes body into a fresh temp file and returns its path.
func writeSource(t *testing.T, name, body string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}

	return path
}

func newFilePage(t *testing.T, cfg page.Config) *page.Page {
	t.Helper()

	p, err := page.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func TestNewReadsFile(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	html := `<div id="a">from disk</div>`
	fromFile := newFilePage(t, page.Config{
		File:   writeSource(t, "index.html", html),
		Width:  320,
		Height: 200,
	})
	fromString := newFilePage(t, page.Config{HTML: html, Width: 320, Height: 200})

	if fromFile.HTML() != html {
		t.Fatalf("HTML = %q", fromFile.HTML())
	}

	if err := fromFile.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if err := fromString.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	boxCenter(t, fromFile, "a")

	if !bytes.Equal(fromFile.PNG(), fromString.PNG()) {
		t.Fatal("file page and string page painted differently")
	}
}

func TestNewReadsThemeFile(t *testing.T) {
	t.Parallel()

	p := newFilePage(t, page.Config{
		HTML:      themePage,
		ThemeFile: writeSource(t, "theme.css", `#swatch{background:#ff0000}`),
		Width:     320,
		Height:    200,
	})

	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if r, g, b := swatchRGB(t, p); r < 1 || g > 0 || b > 0 {
		t.Fatalf("swatch = %v,%v,%v, want red", r, g, b)
	}
}

func TestNewRejectsBothSources(t *testing.T) {
	t.Parallel()

	_, err := page.New(page.Config{
		HTML:   "<p>x</p>",
		File:   writeSource(t, "index.html", "<p>y</p>"),
		Width:  320,
		Height: 200,
	})
	if !errors.Is(err, page.ErrBadSource) {
		t.Fatalf("err = %v", err)
	}
}

func TestNewRejectsMissingFile(t *testing.T) {
	t.Parallel()

	_, err := page.New(page.Config{
		File:   filepath.Join(t.TempDir(), "gone.html"),
		Width:  320,
		Height: 200,
	})
	if !errors.Is(err, page.ErrBadSource) {
		t.Fatalf("err = %v", err)
	}
}

func TestNewBlankFileIsEmptyHTML(t *testing.T) {
	t.Parallel()

	_, err := page.New(page.Config{
		File:   writeSource(t, "blank.html", " \n"),
		Width:  320,
		Height: 200,
	})
	if !errors.Is(err, page.ErrEmptyHTML) {
		t.Fatalf("err = %v", err)
	}
}
