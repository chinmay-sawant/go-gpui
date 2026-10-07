package ui

import (
	"fmt"
	"strings"
)

// buildDialog lays out the CSV import or export overlay.
func (a *App) buildDialog() *Dialog {
	if !a.csv.open {
		return nil
	}

	d := &Dialog{W: 440, H: 310}
	if a.csv.export {
		d.Title = "Export CSV"
		d.Prompt = "Write the active sheet to this path."
		d.Field = "csvsave"
		d.FieldType = "text"
		d.FieldVal = a.csv.path
		d.Buttons = []Box{
			{Action: "csv-write", Text: "Write", W: 88},
			{Action: "csv-cancel", Text: "Cancel", W: 88},
		}
	} else {
		mode := "Replace: off"
		if a.csv.replace {
			mode = "Replace: on"
		}

		d.Title = "Import CSV"
		d.Prompt = "Choose a .csv file, check the preview, then commit."
		d.Field = "csvpath"
		d.FieldType = "file"
		d.FieldVal = a.csv.path
		d.Buttons = []Box{
			{Action: "csv-replace", Text: mode, W: 92},
			{Action: "csv-commit", Text: "Commit", W: 88},
			{Action: "csv-cancel", Text: "Cancel", W: 88},
		}
	}

	if a.csv.ready {
		for i, row := range a.csv.prev.Rows {
			if i == 8 {
				d.Rows = append(d.Rows, "...")
				break
			}

			d.Rows = append(d.Rows, strings.Join(row, " | "))
		}

		d.Note = fmt.Sprintf("%d rows x %d cols", a.csv.prev.Total, a.csv.prev.Cols)
		d.Warnings = a.csv.prev.Warnings
	} else {
		d.Note = a.csv.note
	}

	d.X = a.scrollX + max(16, (a.viewW-d.W)/2)
	d.Y = a.scrollY + max(16, (a.viewH-d.H)/2)
	for i := range d.Buttons {
		d.Buttons[i].H = 26
		d.Buttons[i].X = d.W - (len(d.Buttons)-i)*96
		d.Buttons[i].Y = d.H - 38
		d.Buttons[i].Class = "btn"
	}

	return d
}
