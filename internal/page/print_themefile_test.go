package page

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPDFSourceUsesThemeFile checks a ThemeFile theme reaches Page.PDF.
func TestPDFSourceUsesThemeFile(t *testing.T) {
	t.Parallel()

	theme := `p{color:#ff0000}`
	path := filepath.Join(t.TempDir(), "theme.css")
	if err := os.WriteFile(path, []byte(theme), 0o600); err != nil {
		t.Fatal(err)
	}

	p, err := New(Config{
		HTML:      `<html><head></head><body><p>hi</p></body></html>`,
		ThemeFile: path,
		Width:     640,
		Height:    480,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := p.pdfSource(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(got, theme) {
		t.Fatalf("pdf source misses ThemeFile theme: %q", got)
	}
}
