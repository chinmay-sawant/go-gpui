package parser

import (
	"regexp"
	"time"
)

var (
	reDateTime = regexp.MustCompile(`^\[?(\d{4}-\d{2}-\d{2})([T ])(\d{2}:\d{2}:\d{2})(\.\d+)?(Z|[+-]\d{2}:?\d{2})?`)
	reSlash    = regexp.MustCompile(`^\[?(\d{4}/\d{2}/\d{2}) (\d{2}:\d{2}:\d{2})(\.\d+)?`)
	reUS       = regexp.MustCompile(`^\[?(\d{2}/\d{2}/\d{4}) (\d{2}:\d{2}:\d{2})(\.\d+)?`)
	reSyslog   = regexp.MustCompile(`^([A-Z][a-z]{2} {1,2}\d{1,2} \d{2}:\d{2}:\d{2})`)
	reEpoch    = regexp.MustCompile(`^(\d{9,10})(\.\d+)?\b`)
	reClock    = regexp.MustCompile(`^\[?(\d{2}:\d{2}:\d{2})(\.\d+)?`)
)

// findTime returns the first parseable timestamp in text together with the
// text it matched. A malformed value simply yields ok=false.
func findTime(text string) (time.Time, string, bool) {
	if m := reDateTime.FindStringSubmatch(text); m != nil {
		layout := "2006-01-02" + m[2] + "15:04:05" + fracLayout(m[4]) + zoneLayout(m[5])
		raw := m[1] + m[2] + m[3] + fracText(m[4]) + m[5]
		if t, err := time.Parse(layout, raw); err == nil {
			return t, raw, true
		}
	}

	for _, c := range []struct {
		re     *regexp.Regexp
		layout string
	}{
		{reSlash, "2006/01/02 15:04:05"},
		{reUS, "01/02/2006 15:04:05"},
	} {
		if m := c.re.FindStringSubmatch(text); m != nil {
			raw := m[1] + " " + m[2] + fracText(m[3])
			if t, err := time.Parse(c.layout+fracLayout(m[3]), raw); err == nil {
				return t, raw, true
			}
		}
	}

	return findTimeMore(text)
}
