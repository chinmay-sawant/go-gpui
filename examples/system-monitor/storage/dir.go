package storage

import (
	"os"
	"path/filepath"
)

// DefaultDir returns the per-user data directory, an "ownframe" directory
// holding one subdirectory per example. Calling it does not create anything.
func DefaultDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, "ownframe", "system-monitor"), nil
}
