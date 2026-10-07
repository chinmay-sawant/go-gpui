package storage

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestReadOnlyDirFailsCleanly(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root ignores directory permissions")
	}

	dir := filepath.Join(t.TempDir(), "readonly")

	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}

	if err := os.Chmod(dir, 0o555); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { os.Chmod(dir, 0o755) })

	if _, err := Open(dir); err == nil {
		t.Fatal("Open in a read-only directory succeeded")
	}
}

func TestReadOnlyFileFailsOnWrite(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root ignores file permissions")
	}

	st, dir := openDir(t)

	if _, err := st.CreateWorkbook(context.Background(), "book"); err != nil {
		t.Fatal(err)
	}

	st.Close()

	path := filepath.Join(dir, File)

	if err := os.Chmod(path, 0o444); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { os.Chmod(path, 0o644) })

	st, err := Open(dir)
	if err != nil {
		return // surfaced at open, which is also fine
	}

	defer st.Close()

	if _, err := st.CreateWorkbook(context.Background(), "write"); err == nil {
		t.Fatal("write on a read-only database succeeded")
	}
}

func TestLockedDatabaseSurfacesError(t *testing.T) {
	st, dir := openDir(t)

	other := raw(t, dir)

	if _, err := other.Exec(`BEGIN IMMEDIATE`); err != nil {
		t.Fatal(err)
	}

	defer other.Exec(`ROLLBACK`)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	err := st.SetPref(ctx, "k", "v")
	if err == nil {
		t.Fatal("write on a locked database succeeded")
	}

	if !errors.Is(err, context.DeadlineExceeded) && !strings.Contains(err.Error(), "locked") {
		t.Fatalf("lock error = %v", err)
	}

	if _, err := other.Exec(`ROLLBACK`); err != nil {
		t.Fatal(err)
	}

	if err := st.SetPref(context.Background(), "k", "v"); err != nil {
		t.Fatalf("write after unlock: %v", err)
	}
}
