//go:build !windows

package store

import (
	"os"
	"testing"
)

func TestReadOnlyDirectory(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root bypasses directory permissions")
	}

	dir := t.TempDir()

	st, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}

	if err := st.Close(); err != nil {
		t.Fatal(err)
	}

	if err := os.Chmod(dir, 0o555); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = os.Chmod(dir, 0o755) })

	if _, err := Open(dir); err == nil {
		t.Fatal("opened a database in a read-only directory")
	}
}
