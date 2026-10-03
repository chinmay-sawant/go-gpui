package insights

// streakLevels is the heatmap, one string per weekday, Sunday first.
// Each rune is a Cell level: 0 empty, 1 no dictation, 2..5 teal shades.
var streakLevels = []string{
	"0000001131115555555",
	"0000002111115555555",
	"0000001211115225455",
	"0000012111115535555",
	"0000031111155555555",
	"0000031111155555555",
	"0000023112155555555",
}

// streakCurrent marks the days inside the current streak.
var streakCurrent = []string{
	"............CCCCCCC",
	"............CCCCCCC",
	"............CCCCCCC",
	"............CCCCCCC",
	"...........CCCCCCCC",
	"...........CCCCCCCC",
	"...........CCCCCCCC",
}

// weeks turns the two tables into heatmap cells.
func weeks() [][]Cell {
	rows := make([][]Cell, 0, len(streakLevels))

	for r, levels := range streakLevels {
		row := make([]Cell, 0, len(levels))

		for i := range levels {
			row = append(row, Cell{
				Level:   int(levels[i] - '0'),
				Current: streakCurrent[r][i] == 'C',
			})
		}

		rows = append(rows, row)
	}

	return rows
}
