package page

import (
	"context"
	"time"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"

	"github.com/chinmay-sawant/ownframe/internal/render"
)

// paintFallback paints a bitmap for a page no display list covers. It runs
// after the vector path declines, so a replayable page never pays for it.
func (p *Page) paintFallback(ctx context.Context, styled *css.Document, state render.State, drawStart time.Time, track bool) error {
	ps := time.Now()
	img, boxes, err := render.PaintDocument(ctx, styled, state.Images)
	if track {
		p.stats.paintTime = time.Since(ps)
	}
	if err != nil {
		return err
	}

	p.stats.layouts++
	p.stats.repaints++
	p.stats.lastDraw = time.Since(drawStart)
	p.img = img
	p.display = nil
	p.setBoxes(boxes)
	p.generation++
	p.markFull()
	if track {
		p.recordFull(len(boxes))
	}
	p.applyPending()

	return nil
}
