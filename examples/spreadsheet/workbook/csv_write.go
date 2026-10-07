package workbook

import (
	"encoding/csv"
	"io"
)

// Command aligns the table at at on sheet. Cells past the sheet bounds are
// dropped.
func (t *Table) Command(sheet SheetID, at Pos) Command {
	cmd := Command{Sheet: sheet}

	for i, row := range t.Cells {
		for j, c := range row {
			p := Pos{Row: at.Row + i, Col: at.Col + j}
			if !p.Valid() {
				continue
			}

			cmd.Edits = append(cmd.Edits, CellEdit{Pos: p, Cell: c})
		}
	}

	return cmd
}

// WriteCSV writes the table as CSV. Rows keep their lengths. Cells write
// their raw form, so a formula exports as =source, unless WriteValues is
// set. WriteBOM prefixes the UTF-8 BOM.
func (t *Table) WriteCSV(w io.Writer, opt CSVOptions) error {
	opt = normalCSV(opt)

	if opt.WriteBOM {
		if _, err := w.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
			return err
		}
	}

	cw := csv.NewWriter(w)
	cw.Comma = opt.Delimiter

	for _, row := range t.Cells {
		rec := make([]string, len(row))

		for i, c := range row {
			if opt.WriteValues {
				rec[i] = c.Display()
			} else {
				rec[i] = c.Raw()
			}
		}

		if err := cw.Write(rec); err != nil {
			return err
		}
	}

	cw.Flush()

	return cw.Error()
}

// Snapshot copies a rectangle of a sheet into a table.
func Snapshot(s *Sheet, r Rect) *Table {
	t := &Table{Start: r.Min}

	for row := r.Min.Row; row <= r.Max.Row; row++ {
		cells := make([]Cell, 0, r.Max.Col-r.Min.Col+1)

		for col := r.Min.Col; col <= r.Max.Col; col++ {
			c, _ := s.Cell(Pos{Row: row, Col: col})
			cells = append(cells, c)
		}

		t.Cells = append(t.Cells, cells)
	}

	return t
}
