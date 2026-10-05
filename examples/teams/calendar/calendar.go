// Package calendar is the Calendar menu of the Teams example.
package calendar

// Data is the calendar state the shell prints.
type Data struct {
	Week       int    // index into the weeks
	WeekLabel  string // e.g. "October 5 – October 9"
	Today      int    // column 0..4 that is today, -1 when not in this week
	Compose    bool
	Selected   string
	SelDate    int
	SelStart   int
	SelDur     int
	Days       []Day
	StartTimes []string
	Durs       []string
	Events     []Event

	// allEvents is the whole list the database holds. goWeek filters it
	// by week.
	allEvents []Event
}

// Day is one weekday column header.
type Day struct {
	Label string // "Mon 10/5"
	Today bool
}

// Event is one meeting on the grid.
type Event struct {
	ID       string
	Title    string
	Time     string
	Dur      string
	Color    string
	Location string
	Day      int // column index 0..4
	Selected bool
	Week     int // index into the weeks
}

// DayLabel returns the column label for a day index.
func (d Data) DayLabel(day int) string {
	if day < 0 || day >= len(d.Days) {
		return ""
	}

	return d.Days[day].Label
}
