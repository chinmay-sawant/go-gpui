package page

import "strings"

const formCSS = `<style>[data-gpui-field]{display:inline-block;border:1px solid #c8c2b4;padding:4px 6px;min-width:10em;min-height:1.6em;white-space:pre}[data-gpui-field="textarea"]{white-space:pre-wrap}[data-gpui-field][data-gpui-focus="1"]{border:2px solid #1a56db;padding:3px 5px}[data-gpui-field][data-gpui-selected="1"]{background:#d6e2ff}[data-gpui-field][data-gpui-placeholder="1"]{color:#6b7280}[data-gpui-caret]{display:inline-block;width:1px;height:1em;background:#1c1915}</style>`

// rewriteControls copies source, swapping each kept control for paintable HTML.
func rewriteControls(source string, spans []controlSpan, live map[string]Control, focusID string, selected bool) string {
	if len(spans) == 0 {
		return source
	}

	kept := keepSpans(source, spans)
	cssAt := headOpen(source)
	if cssAt < 0 {
		cssAt = findHead(source)
	}
	if cssAt >= 0 && spanCovers(kept, cssAt) {
		cssAt = -1
	}

	var b strings.Builder
	prev := 0
	placed := false
	write := func(end int) {
		if !placed && cssAt >= prev && cssAt <= end {
			b.WriteString(source[prev:cssAt])
			b.WriteString(formCSS)
			b.WriteString(source[cssAt:end])
			placed = true
			return
		}
		b.WriteString(source[prev:end])
	}
	for _, sp := range kept {
		write(sp.Start)
		ctrl := pickControl(sp, live)
		b.WriteString(replaceControl(source[sp.Start:sp.End], ctrl, focusID, selected))
		prev = sp.End
	}
	write(len(source))
	if !placed {
		return formCSS + b.String()
	}

	return b.String()
}
