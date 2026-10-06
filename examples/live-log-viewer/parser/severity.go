package parser

import (
	"regexp"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

var (
	// A structured field such as level=warn or "severity":"error".
	sevField = regexp.MustCompile(`(?i)"?(level|severity|lvl|levelname|log_level)"?\s*[:=]\s*"?([a-z]+)`)
	// A bracketed level such as [WARN] or (error).
	sevBracket = regexp.MustCompile(`(?i)[\[(<](trace|debug|info|notice|warn|warning|error|err|fatal|critical|crit|panic|alert|emerg)[\])>]`)
	// A level word that opens the line.
	sevLeading = regexp.MustCompile(`(?i)^\s*(trace|debug|info|notice|warn|warning|error|err|fatal|critical|crit|panic|alert|emerg)\b`)
)

// findSeverity looks for a level in a structured field first, then a
// bracketed token, then a leading word. It returns Unknown when the record
// carries no level.
func findSeverity(text string) entry.Severity {
	for _, re := range []*regexp.Regexp{sevField, sevBracket, sevLeading} {
		if m := re.FindStringSubmatch(text); m != nil {
			if s := entry.ParseSeverity(m[len(m)-1]); s != entry.Unknown {
				return s
			}
		}
	}

	return entry.Unknown
}
