package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// formatCount renders an integer with thousands separators.
func formatCount(n int) string {
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}

	var b strings.Builder

	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte(',')
		}

		b.WriteRune(r)
	}

	return b.String()
}

// formatAge renders a duration as "3d 4h", "5h 12m", "1m 30s", or "42s".
func formatAge(d time.Duration) string {
	if d < 0 {
		d = 0
	}

	sec := int(d.Seconds())

	switch {
	case sec < 60:
		return fmt.Sprintf("%ds", sec)
	case sec < 3600:
		return fmt.Sprintf("%dm %ds", sec/60, sec%60)
	case sec < 86400:
		return fmt.Sprintf("%dh %dm", sec/3600, sec%3600/60)
	default:
		return fmt.Sprintf("%dd %dh", sec/86400, sec%86400/3600)
	}
}

// clock formats a sample time for the header.
func clock(t time.Time) string {
	if t.IsZero() {
		return "no sample yet"
	}

	return t.Format("15:04:05")
}

// truncate shortens s to at most n runes, adding an ellipsis.
func truncate(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}

	if n < 2 {
		return string(r[:n])
	}

	return string(r[:n-1]) + "\u2026"
}
