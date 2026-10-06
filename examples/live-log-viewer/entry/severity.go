package entry

import "strings"

// Severity is the level parsed out of one record. Unknown is the zero value
// and sorts below every real level.
type Severity int8

const (
	Unknown Severity = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

var severityNames = [...]string{
	"UNKNOWN", "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL",
}

// String returns the label shown in the UI and in exported files.
func (s Severity) String() string {
	if s < 0 || int(s) >= len(severityNames) {
		return "UNKNOWN"
	}

	return severityNames[s]
}

// ParseSeverity maps a word such as "warning" or "err" to a Severity.
func ParseSeverity(s string) Severity {
	switch strings.ToUpper(strings.Trim(s, "[]()<>: \t\"'")) {
	case "TRACE", "TRC", "VERBOSE":
		return Trace
	case "DEBUG", "DBG", "FINE":
		return Debug
	case "INFO", "INF", "INFORMATION", "NOTICE":
		return Info
	case "WARN", "WRN", "WARNING":
		return Warn
	case "ERROR", "ERR", "SEVERE":
		return Error
	case "FATAL", "CRIT", "CRITICAL", "PANIC", "ALERT", "EMERG":
		return Fatal
	}

	return Unknown
}
