package page_test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

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
