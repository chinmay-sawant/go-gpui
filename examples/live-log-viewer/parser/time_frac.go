package parser

import "strings"

// fracLayout returns the layout piece for a captured fraction, capped at
// nanosecond precision.
func fracLayout(frac string) string {
	if len(frac) < 2 {
		return ""
	}

	n := len(frac) - 1
	if n > 9 {
		n = 9
	}

	return "." + strings.Repeat("0", n)
}

func fracText(frac string) string {
	if len(frac) > 10 {
		return frac[:10]
	}

	return frac
}

func fracNanos(frac string) string {
	frac = strings.TrimPrefix(frac, ".")
	for len(frac) < 9 {
		frac += "0"
	}

	return frac[:9]
}

func zoneLayout(zone string) string {
	switch {
	case zone == "Z", strings.Contains(zone, ":"):
		return "Z07:00"
	case zone != "":
		return "Z0700"
	}

	return ""
}

func parseInt(s string) int64 {
	var n int64

	for _, r := range s {
		if r < '0' || r > '9' {
			return n
		}

		n = n*10 + int64(r-'0')
	}

	return n
}
