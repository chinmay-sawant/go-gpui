package workbook

// TablePage is one keyset page of a table.
type TablePage struct {
	Start Pos
	Rows  [][]Cell
	Next  int
	More  bool
}

// Page returns rows after cursor, at most size of them. Cursor is a row
// index, so unchanged tables page stably.
func (t *Table) Page(cursor, size int) TablePage {
	if size <= 0 {
		size = PageSize
	}

	if cursor < 0 {
		cursor = 0
	}

	if cursor > len(t.Cells) {
		cursor = len(t.Cells)
	}

	end := cursor + size
	if end > len(t.Cells) {
		end = len(t.Cells)
	}

	return TablePage{
		Start: Pos{Row: t.Start.Row + cursor, Col: t.Start.Col},
		Rows:  t.Cells[cursor:end],
		Next:  end,
		More:  end < len(t.Cells),
	}
}
