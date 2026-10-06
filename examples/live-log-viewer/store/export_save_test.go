package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveExportAtomic(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 20})
	if err != nil {
		t.Fatal(err)
	}

	o := ExportOptions{Query: Query{Session: &setup.Session.ID}}
	path := filepath.Join(t.TempDir(), "out.txt")

	if err := st.SaveExport(bg(), o, path); err != nil {
		t.Fatal(err)
	}

	if fi, err := os.Stat(path); err != nil || fi.Size() == 0 {
		t.Fatalf("export file: size=%v err=%v", fi, err)
	}

	bad := filepath.Join(t.TempDir(), "missing", "out.txt")

	if err := st.SaveExport(bg(), o, bad); err == nil {
		t.Fatal("save into a missing directory succeeded")
	}

	if _, err := os.Stat(bad); !os.IsNotExist(err) {
		t.Fatal("a partial export file was left behind")
	}
}
