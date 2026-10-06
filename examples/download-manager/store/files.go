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

// optionalSize returns zero for a missing file and the real size otherwise.
func optionalSize(path string) (int64, error) {
	size, err := fileSize(path)
	if os.IsNotExist(err) {
		return 0, nil
	}

	return size, err
}
