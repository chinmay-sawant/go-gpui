package player

import "strings"

// shellHead opens the document and starts the shared stylesheet.
const shellHead = `<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>Aurora — Audio player</title><style>`

// shellBody opens the content column after the styles.
const shellBody = `</style></head><body><div class="page">`

// shellFoot closes the content column and the document.
const shellFoot = `</div></body></html>`

// styles are concatenated in this order.
var styles = []string{
	"base.css", "sidebar.css", "header.css",
	"now.css", "recent.css", "queue.css",
}

// buildHTML assembles the page from the component fragments.
func buildHTML() string {
	var b strings.Builder

	b.WriteString(shellHead)

	for _, name := range styles {
		b.WriteString(file("components/" + name))
	}

	b.WriteString(shellBody)
	b.WriteString(file("components/sidebar.html"))
	b.WriteString(`<div class="col-main">`)
	b.WriteString(file("components/header.html"))
	b.WriteString(file("components/now.html"))
	b.WriteString(file("components/recent.html"))
	b.WriteString(`</div>`)
	b.WriteString(file("components/queue.html"))
	b.WriteString(shellFoot)

	return b.String()
}
