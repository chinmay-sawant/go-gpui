package page

import (
	"sort"
	"strings"
)

const formCSS = `<style>textarea,select{display:inline-block;border:1px solid #c8c2b4;` +
	`min-width:10em;min-height:1.6em;padding:4px 6px;white-space:pre-wrap}` +
	`textarea[data-gpui-focus="1"],select[data-gpui-focus="1"]{border:2px solid #1a56db}</style>`

// rewriteControls copies source, swapping each kept control for paintable HTML.
func rewriteControls(source string, spans []controlSpan, live map[string]Control, focusID string) string {
	if len(spans) == 0 {
		return source
	}

	kept := keepSpans(source, spans)
	cssAt := findHead(source)
	if cssAt >= 0 && spanCovers(kept, cssAt) {
		cssAt = -1
	}

	var b strings.Builder
	prev := 0
	placed := false
	write := func(end int) {
		if !placed && cssAt >= prev && cssAt < end {
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
		b.WriteString(replaceControl(source[sp.Start:sp.End], ctrl, focusID))
		prev = sp.End
	}
	write(len(source))
	if !placed {
		return formCSS + b.String()
	}

	return b.String()
}

func keepSpans(source string, spans []controlSpan) []controlSpan {
	ordered := make([]controlSpan, len(spans))
	copy(ordered, spans)
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Start < ordered[j].Start
	})

	kept := make([]controlSpan, 0, len(ordered))
	prev := 0
	n := len(source)
	for _, sp := range ordered {
		if sp.Start < prev || sp.End > n || sp.End <= sp.Start {
			continue
		}
		kept = append(kept, sp)
		prev = sp.End
	}

	return kept
}
