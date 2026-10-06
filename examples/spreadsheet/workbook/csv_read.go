package workbook

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

// ParseCSV reads CSV bytes. A UTF-8 BOM is stripped, quoted fields and
// embedded newlines follow encoding/csv rules, and rows may differ in
// length. Empty fields are blank; numeric fields become numbers and a
// leading = becomes a formula, subject to opt.
func ParseCSV(data []byte, opt CSVOptions) (*Table, error) {
	opt = normalCSV(opt)

	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = opt.Delimiter
	r.FieldsPerRecord = -1

	table := &Table{}
	total := 0

	for {
		rec, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		total += len(rec)
		if total > opt.MaxCells {
			return nil, ErrTooLarge
		}

		row := make([]Cell, len(rec))

		for i, f := range rec {
			row[i] = parseField(f, opt)
		}

		table.Cells = append(table.Cells, row)
	}

	return table, nil
}

// normalCSV fills unset option fields.
func normalCSV(opt CSVOptions) CSVOptions {
	if opt.Delimiter == 0 {
		opt.Delimiter = ','
	}

	if opt.MaxCells <= 0 {
		opt.MaxCells = DefaultCSVOptions().MaxCells
	}

	return opt
}

// parseField interprets one field.
func parseField(s string, opt CSVOptions) Cell {
	if s == "" {
		return Cell{}
	}

	if opt.InterpretFormulas && s[0] == '=' {
		return Cell{Kind: Formula, Source: s[1:]}
	}

	if opt.InterpretNumbers {
		if n, ok := plainNumber(strings.TrimSpace(s)); ok {
			return Cell{Kind: Number, Number: n}
		}
	}

	return Cell{Kind: Text, Text: s}
}
