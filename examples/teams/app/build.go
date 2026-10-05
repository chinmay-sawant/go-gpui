package app

import "strings"

// shellHead opens the document and starts the stylesheet.
const shellHead = `<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>Microsoft Teams</title><style>`

// shellBody closes the styles and opens the app row.
const shellBody = `</style></head><body><div class="app">`

// shellFoot closes the document.
const shellFoot = `</div></body></html>`

// buildHTML assembles the app shell from the menu fragments.
func buildHTML() string {
	var b strings.Builder

	b.WriteString(shellHead)
	b.WriteString(styleSheets())
	b.WriteString(shellBody)
	b.WriteString(file("components/rail.html"))
	b.WriteString(`<main class="main">{{if .Note}}<div class="toast">{{.Note}}</div>{{end}}`)

	for i, page := range pages {
		b.WriteString("{{")

		if i > 0 {
			b.WriteString("else ")
		}

		b.WriteString(`if eq .Section "` + page.Name + `"}}`)
		b.WriteString(page.HTML())
	}

	b.WriteString(`{{end}}`)
	b.WriteString(file("components/flyouts.html"))
	b.WriteString(`</main>`)
	b.WriteString(shellFoot)

	return b.String()
}
