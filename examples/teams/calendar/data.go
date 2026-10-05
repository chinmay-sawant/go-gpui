package calendar

// week is one read-only sample week.
type week struct {
	label  string
	today  int
	days   []Day
	events []Event
}

// weeks are the seven sample weeks. Nothing writes to the table.
var weeks = []week{week0, week1, week2, week3, week4, week5, week6}

// defaultWeek is the week that contains today, Monday October 5, 2026.
const defaultWeek = 3

// weekData builds the printable state for one sample week.
// The slices are cloned, so two pages never share mutable state.
func weekData(i int) Data {
	w := weeks[i]

	return Data{
		Week:       i,
		WeekLabel:  w.label,
		Today:      w.today,
		SelDur:     1,
		Days:       append([]Day(nil), w.days...),
		StartTimes: startTimes,
		Durs:       durations,
		Events:     append([]Event(nil), w.events...),
	}
}
