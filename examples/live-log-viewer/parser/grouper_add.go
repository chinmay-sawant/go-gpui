package parser

import (
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Add feeds one raw record and returns the entries it completed.
func (g *Grouper) Add(rec entry.RawRecord) []entry.Entry {
	var out []entry.Entry

	if g.gen != 0 && rec.Generation != g.gen {
		out = g.Flush()
	}

	g.gen = rec.Generation
	parsed := ParseLine(rec.Data)

	if g.pending != nil && isContinuation(parsed.Text) &&
		g.lines < g.opts.MaxLines && g.bytes < g.opts.MaxBytes {
		g.pending.Message += "\n" + parsed.Text
		g.pending.Multiline = true
		g.pending.Bytes += rec.Bytes
		g.pending.Truncated = g.pending.Truncated || rec.Truncated
		g.pending.Malformed = g.pending.Malformed || parsed.Malformed
		g.lines++
		g.bytes += rec.Bytes

		return out
	}

	if g.pending != nil {
		out = append(out, g.emit())
	}

	g.pending = &entry.Entry{
		Position:   rec.Offset,
		Generation: rec.Generation,
		Bytes:      rec.Bytes,
		Time:       parsed.Time,
		TimeRaw:    parsed.TimeRaw,
		TimeOK:     parsed.TimeOK,
		Severity:   parsed.Severity,
		Message:    parsed.Text,
		Truncated:  rec.Truncated,
		Malformed:  parsed.Malformed,
		Partial:    rec.Partial,
	}
	g.lines, g.bytes = 1, rec.Bytes

	if rec.Partial {
		out = append(out, g.emit())
	}

	return out
}

// isContinuation reports whether a record continues the previous one. A
// leading space or tab, or a stack-frame opening, counts.
func isContinuation(text string) bool {
	if text == "" {
		return false
	}

	if text[0] == ' ' || text[0] == '\t' {
		return true
	}

	for _, p := range contPrefixes {
		if strings.HasPrefix(text, p) {
			return true
		}
	}

	return false
}

var contPrefixes = []string{
	"at ", "goroutine ", "Caused by:", "Traceback", "Suppressed:",
	"...", "frame ", "stack=", "panic stack:",
}
