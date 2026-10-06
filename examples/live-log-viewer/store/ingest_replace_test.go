package store

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestIngestFileReplacePartial(t *testing.T) {
	dir := t.TempDir()

	st, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}

	defer st.Close()

	src := fileSource(t, st, dir, "INFO first\nINFO partial")

	in, cancel, done := startFileIngest(t, st, src)

	if got := waitForEntries(t, st, 1, 3*time.Second); got < 1 {
		cancel()
		<-done

		t.Fatalf("entries = %d", got)
	}

	cancel()
	<-done

	if err := in.Close(); err != nil {
		t.Fatal(err)
	}

	logPath := filepath.Join(dir, "app.log")

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := f.WriteString("\n"); err != nil {
		t.Fatal(err)
	}

	f.Close()

	in2, cancel2, done2 := startFileIngest(t, st, src)

	deadline := time.Now().Add(3 * time.Second)

	for time.Now().Before(deadline) {
		p, err := st.LastPosition(bg(), src.ID)
		if err != nil {
			t.Fatal(err)
		}

		if p.OK && !p.Partial {
			break
		}

		time.Sleep(5 * time.Millisecond)
	}

	cancel2()
	<-done2

	if err := in2.Close(); err != nil {
		t.Fatal(err)
	}

	n, err := st.Count(bg(), Query{Source: &src.ID})
	if err != nil {
		t.Fatal(err)
	}

	if n != 2 {
		t.Fatalf("entries = %d, want 2", n)
	}

	last, err := st.LastPosition(bg(), src.ID)
	if err != nil {
		t.Fatal(err)
	}

	if last.Partial {
		t.Fatal("completed line did not replace the partial record")
	}
}
