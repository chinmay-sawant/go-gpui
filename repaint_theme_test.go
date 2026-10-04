package gpui

import (
	"context"
	"testing"
)

const themeRepaintHTML = `<html><body style="margin:0;background:var(--bg,#f4f1ea)">
<div style="width:280px;margin:20px;padding:16px;background:var(--card,#fff)">
<h1 style="color:var(--ink,#1c1915)">{{.Title}}</h1>
<div id="toggle">Switch</div></div></body></html>`

const themeRepaintCSS = `body { background: #102030 } div { background: #203040 } h1 { color: #ffffff }`

// themeCase switches the stylesheet from a click handler.
func themeCase() repaintCase {
	return repaintCase{
		name: "theme",
		newPage: func(t *testing.T) *Page {
			return newPage(t, Config{Title: "Theme", HTML: themeRepaintHTML, Width: 320, Height: 200},
				func(p *Page) {
					p.SetData(map[string]string{"Title": "Live theme"})
					p.Handle(Handlers{Click: func(_ context.Context, box Box) error {
						if box.ID == "toggle" {
							return p.SetTheme(themeRepaintCSS)
						}

						return nil
					}})
				})
		},
		interact: func(t *testing.T, p *Page, ctx context.Context) {
			clickID(t, p, ctx, "toggle")
		},
	}
}
