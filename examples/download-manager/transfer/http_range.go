package transfer

import (
	"fmt"
	"strconv"
	"strings"
)

// parseContentRange splits "bytes 10-19/100" or "bytes */100". A zero
// start with a star span is the 416 shape.
func parseContentRange(value string) (start, end, total int64, err error) {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "bytes ")

	span, totalText, ok := strings.Cut(value, "/")
	if !ok {
		return 0, 0, 0, badRange(value)
	}

	total, err = contentRangeTotal(totalText)
	if err != nil {
		return 0, 0, 0, err
	}

	if span == "*" {
		return 0, 0, total, nil
	}

	left, right, ok := strings.Cut(span, "-")
	if !ok {
		return 0, 0, 0, badRange(value)
	}

	start, err = strconv.ParseInt(left, 10, 64)
	if err != nil || start < 0 {
		return 0, 0, 0, badRange(value)
	}

	end, err = strconv.ParseInt(right, 10, 64)
	if err != nil || end < start {
		return 0, 0, 0, badRange(value)
	}

	return start, end, total, nil
}

// contentRangeTotal parses the total field; "*" and empty mean Unknown.
func contentRangeTotal(text string) (int64, error) {
	text = strings.TrimSpace(text)
	if text == "" || text == "*" {
		return Unknown, nil
	}

	total, err := strconv.ParseInt(text, 10, 64)
	if err != nil || total < 0 {
		return 0, badRange(text)
	}

	return total, nil
}

// badRange wraps a parse failure.
func badRange(value string) error {
	return fmt.Errorf("%w: bad Content-Range %q", ErrBadRange, value)
}
