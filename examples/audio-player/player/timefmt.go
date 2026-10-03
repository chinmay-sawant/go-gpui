package player

import (
	"fmt"
	"strconv"
	"strings"
)

// clock formats seconds as m:ss.
func clock(seconds int) string {
	if seconds < 0 {
		seconds = 0
	}

	return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
}

// trackSeconds reads a m:ss label. The second value reports a real read.
func trackSeconds(length string) (int, bool) {
	parts := strings.Split(length, ":")
	if len(parts) != 2 {
		return 0, false
	}

	minutes, err := strconv.Atoi(parts[0])
	if err != nil || minutes < 0 {
		return 0, false
	}

	seconds, err := strconv.Atoi(parts[1])
	if err != nil || seconds < 0 {
		return 0, false
	}

	return minutes*60 + seconds, true
}

// progressText splits a length at pct into elapsed and remaining labels.
func progressText(length string, pct int) (string, string) {
	total, ok := trackSeconds(length)
	if !ok {
		return "0:00", "-" + length
	}

	elapsed := total * pct / 100

	return clock(elapsed), "-" + clock(total-elapsed)
}

// setProgress moves the seek bar and rewrites both time labels.
func (v *View) setProgress(pct int) {
	v.Progress = clamp(pct, 0, 100)
	v.Elapsed, v.Remaining = progressText(v.Now.Length, v.Progress)
}

func clamp(value, low, high int) int {
	if value < low {
		return low
	}

	if value > high {
		return high
	}

	return value
}
