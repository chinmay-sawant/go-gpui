package store

import (
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// exportLine renders one entry as a single text line, with the timestamp it
// parsed, its severity label, explicit data-quality markers, and the message.
// A multiline message keeps its newlines.
func exportLine(e entry.Entry) string {
	var b strings.Builder

	switch {
	case e.TimeOK && !e.Time.IsZero():
		b.WriteString(e.Time.UTC().Format("2006-01-02T15:04:05.000Z07:00"))
	case e.TimeRaw != "":
		b.WriteString("[" + e.TimeRaw + "]")
	default:
		b.WriteString("-")
	}

	b.WriteString(" ")
	b.WriteString(e.Severity.String())
	b.WriteString(" ")

	for _, flag := range []struct {
		on   bool
		name string
	}{
		{e.Multiline, "multiline"},
		{e.Truncated, "truncated"},
		{e.Malformed, "malformed"},
	} {
		if flag.on {
			b.WriteString("(" + flag.name + ") ")
		}
	}

	b.WriteString(e.Message)
	b.WriteString("\n")

	return b.String()
}
