package fixture

import (
	"net/http"
	"strconv"
	"strings"
)

// writeFull sends the body with a length and a validator.
func writeFull(w http.ResponseWriter, body []byte) {
	w.Header().Set("ETag", etagOK)
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	_, _ = w.Write(body)
}

// parseRange turns "bytes=5-" or "bytes=5-9" into an inclusive span. It
// returns ok=false for anything unsatisfiable.
func parseRange(value string, size int64) (start, end int64, ok bool) {
	value = strings.TrimPrefix(value, "bytes=")
	if value == "" || size == 0 {
		return 0, 0, false
	}

	left, right, found := strings.Cut(value, "-")
	if !found {
		return 0, 0, false
	}

	start, err := strconv.ParseInt(left, 10, 64)
	if err != nil || start < 0 || start >= size {
		return 0, 0, false
	}

	end = size - 1
	if right != "" {
		end, err = strconv.ParseInt(right, 10, 64)
		if err != nil || end < start {
			return 0, 0, false
		}

		if end >= size {
			end = size - 1
		}
	}

	return start, end, true
}
