package ownframe

import (
	"context"
	"testing"
)

const platformRepaintHTML = `<html><body style="margin:0;background:#f4f1ea">
<div style="width:300px;margin:20px;padding:16px;background:#fff">
<h1>Platform</h1><p id="count">{{.Count}}</p>
<button id="inc">Increment</button></div></body></html>`

// platformCase clicks the counter twice.
func platformCase() repaintCase {
	return repaintCase{
		name: "platform",
		newPage: func(t *testing.T) *Page {
			return newPage(t, Config{Title: "Platform", HTML: platformRepaintHTML, Width: 320, Height: 200},
				func(p *Page) {
					count := 0
					p.SetData(map[string]int{"Count": count})
					p.Handle(Handlers{Click: func(_ context.Context, box Box) error {
						if box.ID == "inc" {
							count++
							p.SetData(map[string]int{"Count": count})
						}

						return nil
					}})
				})
		},
		interact: func(t *testing.T, p *Page, ctx context.Context) {
			clickID(t, p, ctx, "inc")
			clickID(t, p, ctx, "inc")
		},
	}
}
