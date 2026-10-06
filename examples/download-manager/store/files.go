package store

import (
	"fmt"
	"os"
)

// fileSize returns a file size and rejects a directory.
func fileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if info.IsDir() {
		return 0, fmt.Errorf("store: %s is a directory", path)
	}

	return info.Size(), nil
}
