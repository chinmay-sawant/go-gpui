package window

import (
	"strings"
	"testing"
)

// TestDevOpsRows checks the Ops tab toggle, counts, and paint list.
func TestDevOpsRows(t *testing.T) {
	t.Parallel()

	s := newDevShell(newDevScreen())
	s.display = devTestDisplay()
	s.dev.ops = true

	rows := s.devOpsRows()
	if len(rows) == 0 {
		t.Fatal("no rows")
	}
	if rows[0].act != devActOpsToggle {
		t.Fatalf("first row act = %v, want devActOpsToggle", rows[0].act)
	}
	if got := rows[0].line.plain(); !strings.HasPrefix(got, "[x]") {
		t.Fatalf("toggle text = %q, want [x]", got)
	}

	var fill, total *devRow
	var list []devRow
	for i := range rows {
		plain := rows[i].line.plain()
		if rows[i].act == devActOp {
			list = append(list, rows[i])
		}
		if rows[i].act == devActNone && strings.Contains(plain, "fill") {
			fill = &rows[i]
		}
		if strings.HasPrefix(plain, "total") {
			total = &rows[i]
		}
	}

	if fill == nil || len(fill.line.spans) == 0 {
		t.Fatalf("no fill count row in %v", rows)
	}
	if sw := fill.line.spans[0]; sw.swatch <= 0 || sw.ink != devFillInk {
		t.Fatalf("fill swatch = %+v, want devFillInk", sw)
	}

	if total == nil || strings.TrimSpace(total.line.plain()) != "total 2" {
		t.Fatalf("total row = %v, want total 2", total)
	}

	if len(list) != 2 || list[0].arg != 0 || list[1].arg != 1 {
		t.Fatalf("list rows = %+v, want args 0 and 1", list)
	}
	if got := list[1].line.plain(); !strings.Contains(got, "hi") {
		t.Fatalf("text row = %q, want hi", got)
	}

	s.display = nil
	rows = s.devOpsRows()
	found := false
	for _, row := range rows {
		found = found || strings.Contains(row.line.plain(), "bitmap fallback")
	}
	if !found {
		t.Fatalf("nil display rows = %v, want bitmap fallback", rows)
	}
}
