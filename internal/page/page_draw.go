package page

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// Redraw fills the template and renders the current size. A page the vector
// replay can draw keeps its display list and no bitmap; any other page keeps
// the rasterized picture from the engine.
func (p *Page) Redraw(ctx context.Context) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	var body strings.Builder
	if err := p.tpl.Execute(&body, p.data); err != nil {
		return err
	}

	p.source = p.syncForm(body.String())
	p.png = nil

	state := p.renderState()

	display, err := render.DisplayListState(ctx, p.source, p.width, p.height, state)
	if err == nil && render.Replayable(display) {
		p.img = nil
		p.display = display
		p.boxes = display.Boxes
		p.generation++

		return nil
	}

	img, boxes, err := render.PaintState(ctx, p.source, p.width, p.height, state)
	if err != nil {
		return err
	}

	p.img = img
	p.display = nil
	p.boxes = boxes
	p.generation++

	return nil
}

// renderState is the input and image state Redraw and PNG share.
func (p *Page) renderState() render.State {
	state := render.State{Hover: p.hover, Active: p.active, Images: p.imageBytes}
	if p.form != nil {
		state.Focus = p.form.focusID
	}

	return state
}
