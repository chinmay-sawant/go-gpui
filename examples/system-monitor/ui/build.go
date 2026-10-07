package ui

import "strings"

// docHead opens the document and starts the base stylesheet.
const docHead = `<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>System Monitor</title><style>`

// docMid closes the styles and opens the app row.
const docMid = `</style></head><body><div class="app">`

// docFoot closes the document.
const docFoot = `</div></body></html>`

// buildHTML assembles the shell, the nav, and the three screen fragments.
// The theme sheet is separate: it arrives through Config.Theme.
func buildHTML() string {
	var b strings.Builder

	b.WriteString(docHead)
	b.WriteString(file("components/base.css"))
	b.WriteString(docMid)
	b.WriteString(file("components/nav.html"))
	b.WriteString(`<main class="main">`)
	b.WriteString(`{{if .Notice}}<div class="notice">{{.Notice}}</div>{{end}}`)
	b.WriteString(`{{if eq .Nav "overview"}}`)
	b.WriteString(file("components/overview.html"))
	b.WriteString(`{{else if eq .Nav "processes"}}`)
	b.WriteString(file("components/processes.html"))
	b.WriteString(`{{else}}`)
	b.WriteString(file("components/detail.html"))
	b.WriteString(`{{end}}</main>`)
	b.WriteString(docFoot)

	return b.String()
}
