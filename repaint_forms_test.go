package ownframe

import (
	"context"
	"testing"
)

const formsRepaintHTML = `<html><head><style>
#agree { width:18px;height:18px }
#agree:checked { outline:3px solid #176b45 }
#email { display:block;width:240px;border:1px solid #999;padding:6px }
</style></head><body style="margin:0;background:#f4f1ea">
<div style="width:280px;margin:20px;padding:16px;background:#fff">
<input id="agree" type="checkbox"><label for="agree">Agree</label>
<input id="email">
</div></body></html>`

// formsCase toggles the checkbox and types into the field.
func formsCase() repaintCase {
	return repaintCase{
		name: "forms",
		newPage: func(t *testing.T) *Page {
			return newPage(t, Config{Title: "Forms", HTML: formsRepaintHTML, Width: 320, Height: 220}, nil)
		},
		interact: func(t *testing.T, p *Page, ctx context.Context) {
			clickID(t, p, ctx, "agree")
			clickID(t, p, ctx, "email")
			typeText(t, p, ctx, "note")
		},
	}
}
