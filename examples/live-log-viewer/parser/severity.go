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
	// A bare level word right after a leading timestamp.
	sevAfterTime = regexp.MustCompile(`(?i)^\[?\d{4}[-/]\d{2}[-/]\d{2}[T ]\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+-]\d{2}:?\d{2})?\]?\s+(trace|debug|info|notice|warn|warning|error|err|fatal|critical|crit|panic|alert|emerg)\b`)
	// A level word that opens the line.
	sevLeading = regexp.MustCompile(`(?i)^\s*(trace|debug|info|notice|warn|warning|error|err|fatal|critical|crit|panic|alert|emerg)\b`)
)

// findSeverity looks for a level in a structured field first, then a
// bracketed token, then the words after or before a leading timestamp. It
// returns Unknown when the record carries no level.
func findSeverity(text string) entry.Severity {
	for _, re := range []*regexp.Regexp{sevField, sevBracket, sevAfterTime, sevLeading} {
		if m := re.FindStringSubmatch(text); m != nil {
			if s := entry.ParseSeverity(m[len(m)-1]); s != entry.Unknown {
				return s
			}
		}
	}

	return entry.Unknown
}
