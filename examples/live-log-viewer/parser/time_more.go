package parser

import "time"

// findTimeMore covers the formats without a full date: syslog, epoch
// seconds, and a bare clock time. A bare time keeps today's date so the
// entry still sorts near the records around it.
func findTimeMore(text string) (time.Time, string, bool) {
	if m := reSyslog.FindStringSubmatch(text); m != nil {
		if t, err := time.Parse("Jan _2 15:04:05", m[1]); err == nil {
			return t.AddDate(time.Now().Year(), 0, 0), m[1], true
		}
	}

	if m := reEpoch.FindStringSubmatch(text); m != nil {
		sec, nsec := parseInt(m[1]), parseInt(fracNanos(m[2]))

		return time.Unix(sec, nsec), m[1] + m[2], true
	}

	if m := reClock.FindStringSubmatch(text); m != nil {
		raw := m[1] + fracText(m[2])
		if t, err := time.Parse("15:04:05"+fracLayout(m[2]), raw); err == nil {
			now := time.Now()

			return time.Date(now.Year(), now.Month(), now.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC), raw, true
		}
	}

	return time.Time{}, "", false
}
