package insights

const (
	// streakWindow is the number of week columns the heatmap shows.
	streakWindow = 19

	// streakCurrentDays is the length of the current streak, in days.
	streakCurrentDays = 52
)

// MaxStreakOffset is the oldest window offset the chevrons can reach.
func MaxStreakOffset() int {
	return len(streakHistory[0]) - streakWindow
}

// BuildStreak returns the card data for a window whose newest column sits
// offset weeks before the end of the history.
func BuildStreak(offset int) Streak {
	offset = min(max(offset, 0), MaxStreakOffset())

	cols := len(streakHistory[0])
	weekdays := len(streakHistory)
	start := cols - streakWindow - offset
	firstCurrent := cols*weekdays - streakCurrentDays

	weeks := make([][]Cell, weekdays)

	for r, row := range streakHistory {
		cells := make([]Cell, streakWindow)

		for i := range cells {
			col := start + i
			cells[i] = Cell{
				Level:   int(row[col] - '0'),
				Current: col*weekdays+r >= firstCurrent,
			}
		}

		weeks[r] = cells
	}

	return Streak{
		Days:    "52",
		Longest: "52",
		Months:  monthsFor(start),
		Weeks:   weeks,
		Prev:    offset < MaxStreakOffset(),
		Next:    offset > 0,
	}
}
