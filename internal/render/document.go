package render

import (
	"context"
	"fmt"
	"image"

	"github.com/chinmay-sawant/blinkless/css"
	"github.com/chinmay-sawant/blinkless/layout"
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

// PaintDocument lays out an already styled document. It is the layout half
// of PaintState, without the parse and the cascade. The image is canvas-sized
// and blank. replay.Picture fills the pixels.
func PaintDocument(
	ctx context.Context,
	styled *css.Document,
	images func(src string) ([]byte, error),
) (image.Image, []layout.Box, error) {
	display, err := DisplayListDocument(ctx, styled, images)
	if err != nil {
		return nil, nil, err
	}

	return canvasImage(display), display.Boxes, nil
}
