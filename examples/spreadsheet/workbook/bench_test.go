package workbook

import (
	"testing"
)

func BenchmarkApplySum(b *testing.B) {
	w := New(1, "bench")
	s := w.AddSheet("S")

	for r := 0; r < 1000; r++ {
		if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
			{Pos: Pos{r, 0}, Cell: ParseInput("1")},
		}}); err != nil {
			b.Fatal(err)
		}
	}

	if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 1}, Cell: ParseInput("=SUM(A1:A1000)")},
	}}); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
			{Pos: Pos{i % 1000, 0}, Cell: ParseInput("2")},
		}}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSeedStress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SeedStress(10000)
	}
}
