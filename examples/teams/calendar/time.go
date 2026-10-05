package calendar

import (
	"fmt"
	"strconv"
	"strings"
)

// minutes turns "9:00 AM" into minutes after midnight.
func minutes(t string) int {
	h, rest, ok := strings.Cut(t, ":")
	if !ok {
		return 0
	}

	m, ap, _ := strings.Cut(rest, " ")
	hi, _ := strconv.Atoi(strings.TrimSpace(h))
	mi, _ := strconv.Atoi(strings.TrimSpace(m))
	if hi == 12 {
		hi = 0
	}

	if ap == "PM" {
		hi += 12
	}

	return hi*60 + mi
}

// nextID returns an event id nothing in the list uses yet.
func nextID(events []Event) string {
	max := 0

	for _, e := range events {
		var n int
		if _, err := fmt.Sscanf(e.ID, "e%d", &n); err == nil && n > max {
			max = n
		}
	}

	return fmt.Sprintf("e%d", max+1)
}

// pickIndex turns one chip action into a clamped option index.
func pickIndex(action, prefix string, count int) int {
	n, err := strconv.Atoi(strings.TrimPrefix(action, prefix))
	if err != nil || n < 0 || n >= count {
		return 0
	}

	return n
}
