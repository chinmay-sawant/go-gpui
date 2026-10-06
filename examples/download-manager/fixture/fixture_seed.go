package fixture

import (
	"net/http"
	"strings"
)

// seedFrom returns a stable seed for a path.
func seedFrom(path string) int {
	sum := 0
	for _, r := range path {
		sum = sum*31 + int(r)
	}

	return sum & 0xffff
}

// requestSeed reads ?seed= for tests that want two different bodies.
func requestSeed(r *http.Request) int {
	text := strings.TrimSpace(r.URL.Query().Get("seed"))
	if text == "" {
		return seedFrom(r.URL.Path)
	}

	seed := 0
	for _, digit := range text {
		if digit < '0' || digit > '9' {
			return seedFrom(r.URL.Path)
		}

		seed = seed*10 + int(digit-'0')
	}

	return seed
}
