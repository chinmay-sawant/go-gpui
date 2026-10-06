package ownframe

import "github.com/chinmay-sawant/ownframe/internal/crash"

// Report writes a local crash file and returns its path.
func Report(title, reason string) (string, error) {
	return crash.Write(title, reason)
}

// SetCrashDir sets the directory Report and a recovered panic use.
func SetCrashDir(dir string) {
	crash.SetDir(dir)
}
