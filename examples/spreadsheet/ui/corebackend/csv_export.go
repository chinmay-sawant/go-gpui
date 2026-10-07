package corebackend

import (
	"bytes"
	"context"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// maxExportCells bounds one CSV export.
const maxExportCells = 500000

// reload replaces the in-memory workbook with the stored one. Undo does
// not survive an import.
func (b *Backend) reload() error {
	wb, err := b.st.LoadWorkbook(context.Background(), b.wb.ID())
	if err != nil {
		return err
	}

	wb.SetHistoryLimit(1)
	b.wb = wb
	b.undo, b.redo = nil, nil

	return nil
}

// ExportCSV writes the used range of the sheet with displayed values. An
// empty sheet exports nothing; a sheet past the bound is refused.
func (b *Backend) ExportCSV(id string) (string, error) {
	sh, ok := b.sheet(id)
	if !ok {
		return "", ErrNoSheet
	}

	rect, ok := sh.UsedRange()
	if !ok {
		return "", nil
	}

	if rect.Count() > maxExportCells {
		return "", ErrExportTooLarge
	}

	opt := workbook.DefaultCSVOptions()
	opt.WriteValues = true

	var buf bytes.Buffer
	if err := workbook.Snapshot(sh, rect).WriteCSV(&buf, opt); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// previewText renders one parsed cell for the dialog.
func previewText(c workbook.Cell) string {
	if c.Kind == workbook.Formula {
		return "=" + c.Source
	}

	return c.Display()
}
