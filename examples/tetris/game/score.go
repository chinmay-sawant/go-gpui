package game

// LineScores are the base points for clearing one to four rows at once.
// A clear scores the base times the level.
var LineScores = [5]int{0, 100, 300, 500, 800}

// lineScore returns the base points for n rows, capping the table.
func lineScore(n int) int {
	if n <= 0 {
		return 0
	}

	if n >= len(LineScores) {
		return LineScores[len(LineScores)-1]
	}

	return LineScores[n]
}

// levelFor returns the level for a line count: one level per ten lines,
// capped at MaxLevel.
func levelFor(lines int) int {
	if lines < 0 {
		lines = 0
	}

	return min(1+lines/10, MaxLevel)
}

// addScore adds n to score and caps the result at MaxScore.
func addScore(score, n int) int {
	if n <= 0 {
		return score
	}

	return min(score+n, MaxScore)
}
