package calendar

// startTimes are the half-hour start slots the meeting card offers.
var startTimes = []string{
	"9:00 AM", "9:30 AM", "10:00 AM", "10:30 AM", "11:00 AM", "11:30 AM",
	"12:00 PM", "12:30 PM", "1:00 PM", "1:30 PM", "2:00 PM", "2:30 PM",
	"3:00 PM", "3:30 PM", "4:00 PM", "4:30 PM", "5:00 PM",
}

// durations are the meeting lengths the card offers. The index of "30 min"
// is the default.
var durations = []string{"15 min", "30 min", "45 min", "60 min"}

// startTime returns the chosen start slot.
func (d Data) startTime() string {
	if d.SelStart < 0 || d.SelStart >= len(d.StartTimes) {
		return "9:00 AM"
	}

	return d.StartTimes[d.SelStart]
}

// duration returns the chosen length.
func (d Data) duration() string {
	if d.SelDur < 0 || d.SelDur >= len(d.Durs) {
		return "30 min"
	}

	return d.Durs[d.SelDur]
}

// dayColumn returns the chosen day, or 0.
func (d Data) dayColumn() int {
	if d.SelDate < 0 || d.SelDate >= len(d.Days) {
		return 0
	}

	return d.SelDate
}
