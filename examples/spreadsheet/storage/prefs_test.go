package storage

import (
	"context"
	"database/sql"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

func TestPrefsRoundTrip(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	if _, found, err := st.GetPref(ctx, "theme"); err != nil || found {
		t.Fatalf("unset pref = found=%v err=%v", found, err)
	}

	if err := st.SetPref(ctx, "theme", "dark"); err != nil {
		t.Fatal(err)
	}

	v, found, err := st.GetPref(ctx, "theme")
	if err != nil || !found || v != "dark" {
		t.Fatalf("pref = %q found=%v err=%v", v, found, err)
	}

	if err := st.SetPref(ctx, "note", "Türkçe 🙂"); err != nil {
		t.Fatal(err)
	}

	v, found, err = st.GetPref(ctx, "note")
	if err != nil || !found || v != "Türkçe 🙂" {
		t.Fatalf("unicode pref = %q found=%v err=%v", v, found, err)
	}
}

func TestTrimRevisions(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)

	for i := 0; i < 10; i++ {
		edit(t, st, w, s, workbook.Pos{Row: i, Col: 0}, "1")
	}

	if got := revisionCount(t, st, w.ID()); got != 10 {
		t.Fatalf("revisions = %d, want 10", got)
	}

	n, err := st.TrimRevisions(ctx, w.ID(), 3, 2)
	if err != nil {
		t.Fatal(err)
	}

	if n != 7 {
		t.Fatalf("trimmed = %d, want 7", n)
	}

	if got := revisionCount(t, st, w.ID()); got != 3 {
		t.Fatalf("revisions after trim = %d, want 3", got)
	}

	if _, err := st.TrimRevisions(ctx, w.ID(), 3, 2); err != nil {
		t.Fatal(err)
	}
}

// revisionCount reads the revision rows through the worker.
func revisionCount(t *testing.T, st *Store, id workbook.ID) int {
	t.Helper()

	var count int

	err := st.do(context.Background(), func(ctx context.Context, db *sql.DB) error {
		return db.QueryRowContext(ctx,
			`SELECT COUNT(*) FROM revisions WHERE workbook_id = ?`, int64(id)).Scan(&count)
	})
	if err != nil {
		t.Fatal(err)
	}

	return count
}
