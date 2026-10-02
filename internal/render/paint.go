// Package render paints the screen image from layout.Lay and reads that same
// placement as a display list from layout.DisplayList. Paint flattens the
// placement to one picture. DisplayList keeps it as vector operations. It
// does not build a PDF.
package render

import (
	"context"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/html"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// Paint parses source and returns the screen picture and hit boxes.
func Paint(ctx context.Context, source string, width, height int) (image.Image, []layout.Box, error) {
	return PaintState(ctx, source, width, height, State{})
}

// PaintState is Paint with the runtime pointer and focus state.
func PaintState(ctx context.Context, source string, width, height int, state State) (image.Image, []layout.Box, error) {
	doc, err := html.Parse([]byte(source))
	if err != nil {
		return nil, nil, err
	}

	styled, err := css.Apply(ctx, doc, state.options(width, height))
	if err != nil {
		return nil, nil, err
	}

	placed, err := layout.Lay(ctx, styled)
	if err != nil {
		return nil, nil, err
	}

	return placed.Image(), placed.Boxes(), nil
}
