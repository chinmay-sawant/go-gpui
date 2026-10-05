package calendar

// week4 is the sample week of October 12, 2026.
var week4 = week{
	label: "October 12 – October 16",
	today: -1,
	days: []Day{
		{"Mon 10/12", false},
		{"Tue 10/13", false},
		{"Wed 10/14", false},
		{"Thu 10/15", false},
		{"Fri 10/16", false},
	},
	events: []Event{
		{"e9", "Press and gala prep with Pepper Potts", "9:00 AM", "45 min", "violet", "Avengers Tower 42", 0, false},
		{"e10", "Phase 6 roadmap review", "11:00 AM", "60 min", "blue", "Stark Industries R&D", 0, false},
		{"e11", "Suit design workshop with Shuri", "10:00 AM", "90 min", "orange", "Wakanda Lab 3", 1, false},
		{"e12", "1:1 with James Rhodes", "9:30 AM", "30 min", "green", "Avengers Compound", 2, false},
		{"e13", "Recruiting panel - Kate Bishop", "2:00 PM", "45 min", "red", "Avengers Tower 42", 2, false},
		{"e14", "S.H.I.E.L.D. debrief", "11:30 AM", "45 min", "red", "S.H.I.E.L.D. Helicarrier", 3, false},
		{"e15", "Suit demo Friday", "3:30 PM", "45 min", "blue", "Stark Industries R&D", 4, false},
	},
}
