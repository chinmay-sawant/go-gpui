package app

import "strings"

// shellHead opens the document and starts the stylesheet.
const shellHead = `<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>Wispr Flow</title><style>`

// shellBody closes the styles and opens the app row.
const shellBody = `</style></head><body><div class="app">`

// shellFoot closes the document.
const shellFoot = `</body></html>`

// buildHTML assembles the app shell from the page fragments.
func buildHTML() string {
	var b strings.Builder

	b.WriteString(shellHead)
	b.WriteString(styleSheets())
	b.WriteString(shellBody)
	b.WriteString(file("components/sidebar.html"))
	b.WriteString(`<main class="main">{{if .Note}}<div class="toast">{{.Note}}</div>{{end}}`)

	for i, page := range pages {
		b.WriteString("{{")

		if i > 0 {
			b.WriteString("else ")
		}

		b.WriteString(`if eq .ActivePage "` + page.Name + `"}}`)
		b.WriteString(page.HTML())
	}

	b.WriteString(`{{end}}`)
	b.WriteString(`</main></div>`)
	b.WriteString(`{{if .GetApp}}`)
	b.WriteString(file("components/getapp.html"))
	b.WriteString(`{{end}}`)
	b.WriteString(shellFoot)

	return b.String()
}
