// Package parser turns raw records into entries. It extracts severity and a
// timestamp where it can, tolerates malformed values, replaces invalid UTF-8
// with U+FFFD, and groups continuation lines such as stack frames into one
// multiline entry.
package parser

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Parsed is what one primary record contributes before grouping.
type Parsed struct {
	Text      string
	Time      time.Time
	TimeRaw   string
	TimeOK    bool
	Severity  entry.Severity
	Malformed bool
}

// ParseLine reads one record body. A trailing CR is dropped; invalid UTF-8
// is replaced with U+FFFD and reported through Parsed.Malformed.
func ParseLine(data []byte) Parsed {
	text := strings.TrimSuffix(string(data), "\r")

	p := Parsed{Text: text}
	if !utf8.ValidString(text) {
		p.Text = strings.ToValidUTF8(text, "\uFFFD")
		p.Malformed = true
	}

	p.Time, p.TimeRaw, p.TimeOK = findTime(p.Text)
	p.Severity = findSeverity(p.Text)

	return p
}
