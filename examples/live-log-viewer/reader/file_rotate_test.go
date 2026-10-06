package reader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileRotation(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "app.log")

	writeLog(t, path, "old\n")

	r := testFile(t, path, 0)
	recs := drain(t, r, 1)

	if string(recs[0].Data) != "old" || recs[0].Generation != 1 {
		t.Fatalf("first = %+v", recs[0])
	}

	if err := os.Rename(path, path+".1"); err != nil {
		t.Fatal(err)
	}

	writeLog(t, path, "new\n")

	more := drain(t, r, 1)
	got := more[0]

	if got.Generation != 2 || got.Offset != 0 || string(got.Data) != "new" {
		t.Fatalf("rotated record = %+v", got)
	}
}

func TestFileTruncate(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "first line\n")

	r := testFile(t, path, 0)
	drain(t, r, 1)

	writeLog(t, path, "b\n")

	more := drain(t, r, 1)
	got := more[0]

	if got.Generation != 2 || got.Offset != 0 || string(got.Data) != "b" {
		t.Fatalf("truncated record = %+v", got)
	}
}

func TestFileDeleteRecreate(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "gone\n")

	r := testFile(t, path, 0)
	drain(t, r, 1)

	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}

	writeLog(t, path, "back\n")

	more := drain(t, r, 1)
	got := more[0]

	if got.Generation != 2 || got.Offset != 0 || string(got.Data) != "back" {
		t.Fatalf("recreated record = %+v", got)
	}
}

func TestFileMissingThenFound(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "later.log")

	r := testFile(t, path, 0)

	writeLog(t, path, "appeared\n")

	recs := drain(t, r, 1)
	if string(recs[0].Data) != "appeared" {
		t.Fatalf("record = %+v", recs[0])
	}
}
