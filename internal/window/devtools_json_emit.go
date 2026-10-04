package window

import "strings"

// devJSONPrinter accumulates the rows for one document.
type devJSONPrinter struct {
	rows      []devRow
	collapsed map[string]bool
}

// emit appends the rows for one node. depth is the nesting level, lead is
// the member prefix, tail is a trailing comma, and root keeps the top node
// unfoldable.
func (p *devJSONPrinter) emit(n *devJSONNode, depth int, lead devLine, tail string, root bool) {
	indent := strings.Repeat("  ", depth)

	if n.kind == 's' {
		p.rows = append(p.rows, devRow{line: devJSONLine(indent, lead, n.scalar, n.ink, tail)})

		return
	}

	open, close := "{", "}"
	if n.kind == 'a' {
		open, close = "[", "]"
	}

	if len(n.kids) == 0 {
		p.rows = append(p.rows, devRow{line: devJSONLine(indent, lead, open+close, devPunctInk, tail)})

		return
	}

	if !root && p.collapsed[n.path] {
		p.rows = append(p.rows, devRow{
			line: devJSONLine(indent, lead, open+"..."+close, devPunctInk, tail),
			act:  devActJSON,
			key:  n.path,
		})

		return
	}

	row := devRow{line: devJSONLine(indent, lead, open, devPunctInk, "")}
	if !root {
		row.act, row.key = devActJSON, n.path
	}

	p.rows = append(p.rows, row)

	for i, kid := range n.kids {
		kidLead, kidTail := devLine{}, ","

		if i == len(n.kids)-1 {
			kidTail = ""
		}

		if n.kind == 'o' {
			kidLead = devJSONLead(kid.key)
		}

		p.emit(kid, depth+1, kidLead, kidTail, false)
	}

	p.rows = append(p.rows, devRow{line: devJSONPlain(indent, close, tail)})
}
