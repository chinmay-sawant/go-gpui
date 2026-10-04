package window

import (
	"fmt"
	"time"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// devMS formats a duration for the panel. A zero duration reads as zero
// instead of a negative or rounded value.
func devMS(d time.Duration) string {
	if d <= 0 {
		return "0.0ms"
	}

	return fmt.Sprintf("%.1fms", float64(d)/float64(time.Millisecond))
}

// devErr prints an empty error as a dash for a counter row.
func devErr(err string) string {
	if err == "" {
		return "-"
	}

	return err
}

// devClip shortens a string to n runes, with dots when it was cut.
func devClip(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}

	return string(r[:n]) + "..."
}

// devBoxName is tag or tag#id for labels and the target row.
func devBoxName(box layout.Box) string {
	if box.ID != "" {
		return box.Tag + "#" + box.ID
	}

	return box.Tag
}
