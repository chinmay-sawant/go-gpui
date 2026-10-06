package workbook

import (
	"math"
	"strconv"
	"strings"
)

// plainNumber parses the number forms a spreadsheet accepts. Go's
// ParseFloat also takes hex floats such as 0x10, which typed text should
// keep as text.
func plainNumber(s string) (float64, bool) {
	if strings.ContainsAny(s, "xXpP") {
		return 0, false
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil || math.IsNaN(n) || math.IsInf(n, 0) {
		return 0, false
	}

	return n, true
}
