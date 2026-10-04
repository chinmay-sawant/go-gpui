package gpui

import (
	"context"
	"testing"
)

const scrollRepaintHTML = `<html><body style="margin:0;background:#f4f1ea">
<div style="height:800px"></div>
<div style="width:300px;margin:20px;padding:16px;background:#fff">
<p id="status">{{.Status}}</p><button id="far">Far</button></div></body></html>`

// scrollCase clicks a control below the fold of a page taller than its frame.
func scrollCase() repaintCase {
	return repaintCase{
		name: "scrolling",
		newPage: func(t *testing.T) *Page {
			return newPage(t, Config{Title: "Scrolling", HTML: scrollRepaintHTML, Width: 320, Height: 200},
				func(p *Page) {
					p.SetData(map[string]string{"Status": ""})
					p.Handle(Handlers{Click: func(_ context.Context, box Box) error {
						if box.ID == "far" {
							p.SetData(map[string]string{"Status": "reached"})
						}

						return nil
					}})
				})
		},
		interact: func(t *testing.T, p *Page, ctx context.Context) {
			clickID(t, p, ctx, "far")
		},
	}
}
