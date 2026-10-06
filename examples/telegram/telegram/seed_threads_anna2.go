package telegram

// seedAnnaToday is Anna's morning planning, right before the dinner plan.
func seedAnnaToday() []Message {
	return []Message{
		{ID: "anna-b01", Text: "Morning! The rain finally stopped.", Time: "08:41"},
		{ID: "anna-b02", Text: "About time. I'm in meetings until ten.", Time: "08:44", Own: true, Read: true},
		{ID: "anna-b03", Text: "Good luck. Text me when you're out.", Time: "08:47"},
		{ID: "anna-b04", Text: "Out. Coffee?", Time: "10:12", Own: true, Read: true},
		{ID: "anna-b05", Text: "Can't, stuck at the office until lunch.", Time: "10:15"},
		{ID: "anna-b06", Text: "Then dinner later. Somewhere by the river?", Time: "10:18", Own: true, Read: true},
		{ID: "anna-b07", Text: "Yes! That new place opened last week.", Time: "10:21"},
		{ID: "anna-b08", Text: "The one with the small menu?", Time: "10:24", Own: true, Read: true},
		{ID: "anna-b09", Text: "That's the one. I don't know if it takes bookings.", Time: "10:28"},
		{ID: "anna-b10", Text: "I can call them if you want.", Time: "10:31", Own: true, Read: true},
		{ID: "anna-b11", Text: "Please do. I'll finish this report meanwhile.", Time: "10:35"},
		{ID: "anna-b12", Text: "No promises, but I'll try.", Time: "10:38", Own: true, Read: true},
		{ID: "anna-b13", Text: "Ha! Just text me the time.", Time: "10:41"},
		{ID: "anna-b14", Text: "Report done. Freedom!", Time: "10:44"},
		{ID: "anna-b15", Text: "Talk before lunch.", Time: "10:46"},
	}
}
