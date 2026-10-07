package workbook

import (
	"fmt"
)

// fillText fills a sheet of Unicode, quoted text, embedded newlines, and
// text that only looks numeric.
func fillText(w *Workbook, s *Sheet) {
	for r := 0; r < DummyRows; r++ {
		row := r + 1

		w.put(s, Pos{r, 0}, Cell{Kind: Text, Text: demoLabels[r%len(demoLabels)]})
		w.put(s, Pos{r, 1}, Cell{Kind: Text, Text: `He said "hi", twice`})

		if r%5 == 0 {
			w.put(s, Pos{r, 2}, Cell{Kind: Text, Text: "line one\nline two"})
		}

		w.put(s, Pos{r, 3}, formulaCell("A%d", row))
		w.put(s, Pos{r, 4}, Cell{Kind: Text, Text: fmt.Sprintf("%03d", r)})
		w.put(s, Pos{r, 5}, Cell{Kind: Text, Text: "0"})

		if r%4 == 0 {
			w.put(s, Pos{r, 8}, Cell{Kind: Number, Number: 1})
			w.put(s, Pos{r, 9}, formulaCell("I%d*2", row))
		}
	}
}
