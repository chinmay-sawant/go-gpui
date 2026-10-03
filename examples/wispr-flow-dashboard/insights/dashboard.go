package insights

import "strings"

// shellHead opens the document and starts the shared stylesheet.
const shellHead = `<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>Insights</title><style>`

// shellBody opens the content column after the styles.
const shellBody = `</style></head><body><div class="page">`

// shellFoot closes the content column and the document.
const shellFoot = `</div></body></html>`

// styles are concatenated in this order.
var styles = []string{
	"base.css", "header.css", "tabs.css", "wpm.css",
	"fixes.css", "words.css", "apps.css", "streak.css",
}

// buildHTML assembles the page from the component fragments.
func buildHTML() string {
	var b strings.Builder

	b.WriteString(shellHead)

	for _, name := range styles {
		b.WriteString(file("components/" + name))
	}

	b.WriteString(shellBody)
	b.WriteString(file("components/header.html"))
	b.WriteString(file("components/tabs.html"))
	b.WriteString(`{{if eq .ActiveTab "usage"}}<div class="top-grid">`)
	b.WriteString(file("components/wpm.html"))
	b.WriteString(file("components/fixes.html"))
	b.WriteString(file("components/words.html"))
	b.WriteString(`</div><div class="bottom-grid">`)
	b.WriteString(file("components/apps.html"))
	b.WriteString(file("components/streak.html"))
	b.WriteString(`</div>{{else}}`)
	b.WriteString(file("components/empty.html"))
	b.WriteString(`{{end}}`)
	b.WriteString(shellFoot)

	return b.String()
}
