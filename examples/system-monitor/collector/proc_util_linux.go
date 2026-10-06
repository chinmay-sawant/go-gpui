//go:build linux

package collector

import "os"

// pathExists reports whether a path is present. Probing with Stat is enough;
// the contents decide whether a value is available.
func pathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// isDigits reports whether s is a non-empty run of ASCII digits.
func isDigits(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}
