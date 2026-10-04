// Package cat displays the generated backpack cats in a transparent page.
package cat

import (
	_ "embed"
	"fmt"
	"io/fs"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/desktop-cat/assets"
)

//go:embed cat.html
var source string

// New returns a page that cycles through the backpack-cat collection.
func New() (*gpui.Page, error) {
	return NewVariant(0)
}

// NewVariant selects a one-based expression. Zero cycles through every cat.
func NewVariant(variant int) (*gpui.Page, error) {
	files, err := fs.Glob(assets.Cats, "cat_images/*.png")
	if err != nil {
		return nil, err
	}

	if variant < 0 || variant > len(files) || len(files) == 0 {
		return nil, fmt.Errorf("desktop cat: variant must be 0 through %d", len(files))
	}

	companion, err := newCompanion(variant, files, NewInbox())
	if err != nil {
		return nil, err
	}
	return companion.Page, nil
}
