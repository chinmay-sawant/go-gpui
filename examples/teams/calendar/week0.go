package calendar

// week0 is the sample week of September 14, 2026.
var week0 = week{
	label: "September 14 – September 18",
	today: -1,
	days: []Day{
		{"Mon 9/14", false},
		{"Tue 9/15", false},
		{"Wed 9/16", false},
		{"Thu 9/17", false},
		{"Fri 9/18", false},
	},
	events: []Event{
		{"e16", "Avengers daily standup", "9:00 AM", "15 min", "blue", "Avengers Tower 42", 0, false},
		{"e17", "Mark 85 suit diagnostics", "10:30 AM", "45 min", "violet", "Stark Industries R&D", 0, false},
		{"e18", "1:1 with Bruce Banner", "1:00 PM", "30 min", "green", "Avengers Compound", 1, false},
		{"e19", "Wakanda vibranium sync", "2:30 PM", "45 min", "orange", "Wakanda Lab 3", 2, false},
		{"e20", "Briefing with Nick Fury", "11:00 AM", "60 min", "violet", "S.H.I.E.L.D. Helicarrier", 3, false},
		{"e21", "Web shooters with Peter Parker", "3:30 PM", "45 min", "blue", "Queens", 4, false},
	},
}
