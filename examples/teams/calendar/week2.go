package calendar

// week2 is the sample week of September 28, 2026.
var week2 = week{
	label: "September 28 – October 2",
	today: -1,
	days: []Day{
		{"Mon 9/28", false},
		{"Tue 9/29", false},
		{"Wed 9/30", false},
		{"Thu 10/1", false},
		{"Fri 10/2", false},
	},
	events: []Event{
		{"e28", "Avengers daily standup", "9:00 AM", "15 min", "blue", "Avengers Tower 42", 0, false},
		{"e29", "Mission briefing with Maria Hill", "11:00 AM", "60 min", "violet", "S.H.I.E.L.D. Helicarrier", 0, false},
		{"e30", "Wanda and the Westview report", "9:30 AM", "45 min", "orange", "Sanctum Sanctorum", 1, false},
		{"e31", "1:1 with Natasha Romanoff", "1:30 PM", "30 min", "green", "Avengers Compound", 2, false},
		{"e32", "Vibranium allocation review", "2:00 PM", "90 min", "red", "Wakanda Lab 3", 3, false},
		{"e33", "S.H.I.E.L.D. debrief", "3:00 PM", "45 min", "green", "S.H.I.E.L.D. Helicarrier", 4, false},
	},
}
