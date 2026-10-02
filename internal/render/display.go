package render

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/html"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// DisplayOp is one vector operation from the placement. It is the layout
// package's operation type under a local name, so callers of this package do
// not import the engine to name it. Read the rare payload through its
// accessor methods (LinkURI, ImageBytes, Transform, Opacity, and the rest);
// the zero value carries no payload and a promoted field read on it panics.
type DisplayOp = layout.DisplayOp

// Display is a retained display list plus the canvas it was laid out for.
type Display = layout.Display

// DisplayList parses source and returns the placement as vector operations
// instead of as a picture. width and height are the frame size in CSS pixels,
// the same arguments Paint takes.
//
// The operations carry geometry, color, and font handles, so a caller can draw
// them onto its own canvas and keep text as glyphs rather than as pixels. Paint
// stays the entry for anything that wants the picture, including the web host's
// GET /frame.png.
func DisplayList(ctx context.Context, source string, width, height int) (*Display, error) {
	return DisplayListState(ctx, source, width, height, State{})
}

// DisplayListState is DisplayList with the runtime pointer and focus state.
func DisplayListState(ctx context.Context, source string, width, height int, state State) (*Display, error) {
	doc, err := html.Parse([]byte(source))
	if err != nil {
		return nil, fmt.Errorf("render: parse: %w", err)
	}

	styled, err := css.Apply(ctx, doc, state.options(width, height))
	if err != nil {
		return nil, fmt.Errorf("render: css: %w", err)
	}

	display, err := layout.DisplayList(ctx, styled)
	if err != nil {
		return nil, fmt.Errorf("render: display: %w", err)
	}

	return display, nil
}
