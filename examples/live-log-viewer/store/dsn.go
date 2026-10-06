package store

import (
	"net/url"
	"os"
	"path/filepath"
)

// DefaultDir returns the per-user data directory for this example:
// os.UserConfigDir()/ownframe/live-log-viewer.
func DefaultDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, "ownframe", "live-log-viewer"), nil
}

// dsn builds a modernc.org/sqlite URI. The path is made absolute and turned
// into a file URL, so spaces, Unicode, Windows drive letters, and URI
// punctuation all survive. Busy timeout, foreign keys, and synchronous are
// applied to the physical connection through _pragma parameters.
func dsn(path string) string {
	if path == Memory {
		return sqliteMemory
	}

	if abs, err := filepath.Abs(path); err == nil {
		path = abs
	}

	p := filepath.ToSlash(path)
	if len(p) >= 2 && p[1] == ':' {
		p = "/" + p
	}

	return (&url.URL{Scheme: "file", Path: p}).String() + sqliteExtra
}
