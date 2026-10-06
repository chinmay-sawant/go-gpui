package storage

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// TestCloseWithWriteInFlight checks that Close and an in-flight save never
// panic and that the caller hears either the result or ErrClosed.
func TestCloseWithWriteInFlight(t *testing.T) {
	st, _ := openDir(t)
	w, s := newBook(t, st)

	cmd := []workbook.CellEdit{{Pos: workbook.Pos{Row: 0, Col: 0}, Cell: workbook.ParseInput("x")}}

	results := make(chan error, 1)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		_, err := st.SaveCells(context.Background(), w.ID(), s.ID(), cmd, 0, "inflight")
		results <- err
	}()

	if err := st.Close(); err != nil {
		t.Fatal(err)
	}

	wg.Wait()
	close(results)

	for err := range results {
		if err != nil && !errors.Is(err, ErrClosed) {
			t.Fatalf("in-flight save: %v", err)
		}
	}

	if err := st.Close(); err != nil {
		t.Fatal("second close:", err)
	}
}

func TestCheckpointTruncatesWAL(t *testing.T) {
	st, dir := openDir(t)
	w, s := newBook(t, st)

	edit(t, st, w, s, workbook.Pos{Row: 0, Col: 0}, "1")

	if err := st.Checkpoint(context.Background()); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(filepath.Join(dir, File+"-wal"))
	if err != nil {
		t.Fatal(err)
	}

	if info.Size() != 0 {
		t.Fatalf("wal size = %d after TRUNCATE checkpoint", info.Size())
	}
}
