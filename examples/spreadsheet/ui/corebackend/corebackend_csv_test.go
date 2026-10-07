package corebackend

import (
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/ui"
)

func TestCSVRoundTrip(t *testing.T) {
	b := newBackend(t, t.TempDir())
	id := sheetID(t, b, "Numbers")

	prev, err := b.PreviewCSV(id, []byte("name,qty\nbolt,3\n"), true)
	if err != nil {
		t.Fatal(err)
	}

	if prev.Total != 2 || prev.Cols != 2 || prev.Rows[1][1] != "3" {
		t.Fatalf("preview = %+v", prev)
	}

	res, err := b.CommitCSV(id, []byte("name,qty\nbolt,3\n"), true)
	if err != nil {
		t.Fatal(err)
	}

	if res.Rows != 2 || res.Cols != 2 {
		t.Fatalf("import = %+v", res)
	}

	cells, err := b.Range(id, ui.Area{R0: 0, C0: 0, R1: 1, C1: 1})
	if err != nil {
		t.Fatal(err)
	}

	if cells[0].Raw != "name" || cells[3].Raw != "3" || !cells[3].Num {
		t.Fatalf("imported = %+v", cells)
	}

	out, err := b.ExportCSV(id)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(out, "name,qty\n") || !strings.Contains(out, "bolt,3") {
		t.Fatalf("export = %q", out)
	}
}
