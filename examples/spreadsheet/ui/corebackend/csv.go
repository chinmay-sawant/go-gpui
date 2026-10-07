package corebackend

import (
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/ui"
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// previewRows is how many rows the import dialog shows.
const previewRows = 8

// PreviewCSV parses the bytes and returns the first rows without touching
// the workbook. The whole parse is bounded by the core's MaxCells.
func (b *Backend) PreviewCSV(id string, data []byte, replace bool) (ui.Preview, error) {
	if _, ok := b.sheet(id); !ok {
		return ui.Preview{}, ErrNoSheet
	}

	table, err := workbook.ParseCSV(data, workbook.DefaultCSVOptions())
	if err != nil {
		return ui.Preview{}, err
	}

	page := table.Page(0, previewRows)
	rows := make([][]string, 0, len(page.Rows))
	short := false

	for _, row := range page.Rows {
		if len(row) != table.Cols() {
			short = true
		}

		rec := make([]string, 0, len(row))
		for _, c := range row {
			rec = append(rec, previewText(c))
		}

		rows = append(rows, rec)
	}

	prev := ui.Preview{Rows: rows, Total: table.Rows(), Cols: table.Cols()}
	if page.More {
		prev.Warnings = append(prev.Warnings,
			fmt.Sprintf("preview shows %d of %d rows", len(rows), table.Rows()))
	}

	if short {
		prev.Warnings = append(prev.Warnings, "some rows are shorter than the widest row")
	}

	return prev, nil
}
