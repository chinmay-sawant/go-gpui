package page

import (
	"context"
	"strings"
	"time"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// Redraw fills the template and renders the current size. A page the vector
// replay can draw keeps its display list and no bitmap; any other page keeps
// the rasterized picture from the engine. The parsed tree and its sheets are
// reused until the executed source or the theme changes.
func (p *Page) Redraw(ctx context.Context) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	start := time.Now()
	p.stats.redraws++
	defer func() { p.stats.lastRedraw = time.Since(start) }()

	var body strings.Builder
	if err := p.tpl.Execute(&body, p.data); err != nil {
		return err
	}

	p.source = p.syncForm(body.String())
	p.png = nil

	state := p.renderState()

	styled, err := p.styledDocument(ctx, p.source, state)
	if err != nil {
		return err
	}

	drawStart := time.Now()

	display, derr := render.DisplayListDocument(ctx, styled, state.Images)
	if derr == nil && render.Replayable(display) {
		p.stats.layouts++
		p.stats.repaints++
		p.stats.lastDraw = time.Since(drawStart)
		p.img = nil
		p.display = display
		p.boxes = display.Boxes
		p.generation++

		return nil
	}

	img, boxes, err := render.PaintDocument(ctx, styled, state.Images)
	if err != nil {
		return err
	}

	p.stats.layouts++
	p.stats.repaints++
	p.stats.lastDraw = time.Since(drawStart)
	p.img = img
	p.display = nil
	p.boxes = boxes
	p.generation++

	return nil
}

// renderState is the input, theme, and image state Redraw and PNG share.
func (p *Page) renderState() render.State {
	state := render.State{Hover: p.hover, Active: p.active, Theme: p.theme, Images: p.imageBytes}
	if p.form != nil {
		state.Focus = p.form.focusID
	}

	return state
}
