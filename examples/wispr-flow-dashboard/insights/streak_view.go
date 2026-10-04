package insights

const (
	// streakWindow is the number of week columns the heatmap shows.
	streakWindow = 19

	// streakCurrentDays is the length of the current streak, in days.
	streakCurrentDays = 52
)

// maxStreakOffset is the oldest window offset the chevrons can reach.
func maxStreakOffset() int {
	return len(streakHistory[0]) - streakWindow
}

// shiftStreak scrolls the heatmap by step weeks; positive steps into the
// past. The offset stops at the oldest and newest windows.
func (a *App) shiftStreak(step int) {
	a.streak = min(max(a.streak+step, 0), maxStreakOffset())
	a.view.Streak = buildStreak(a.streak)
}

// buildStreak returns the card data for a window whose newest column sits
// offset weeks before the end of the history.
func buildStreak(offset int) Streak {
	offset = min(max(offset, 0), maxStreakOffset())

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
		Prev:    offset < maxStreakOffset(),
		Next:    offset > 0,
	}
}
