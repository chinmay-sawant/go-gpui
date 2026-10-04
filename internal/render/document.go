package render

import (
	"context"
	"fmt"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// DisplayListDocument turns an already styled document into vector
// operations. It is the layout half of DisplayListState, without the parse
// and the cascade.
func DisplayListDocument(
	ctx context.Context,
	styled *css.Document,
	images func(src string) ([]byte, error),
) (*Display, error) {
	display, err := layout.DisplayListOptions(ctx, styled, layout.Options{Images: images})
	if err != nil {
		return nil, fmt.Errorf("render: display: %w", err)
	}

	return display, nil
}

// PaintDocument rasterizes an already styled document. It is the layout half
// of PaintState, without the parse and the cascade.
func PaintDocument(
	ctx context.Context,
	styled *css.Document,
	images func(src string) ([]byte, error),
) (image.Image, []layout.Box, error) {
	placed, err := layout.LayOptions(ctx, styled, layout.Options{Images: images})
	if err != nil {
		return nil, nil, fmt.Errorf("render: paint: %w", err)
	}

	return placed.Image(), placed.Boxes(), nil
}
